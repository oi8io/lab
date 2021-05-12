package biu

import (
	"fmt"
	"log"
	"oi.io/apps/biu/biu/singleflight"
	"sync"
)

var (
	mu     sync.Mutex
	groups = make(map[string]*CacheGroup)
)

//CacheGroup
type CacheGroup struct {
	name         string
	mainCache    iCache
	sourceGetter SourceGetter
	peers        PeerPicker
	loader       *singleflight.Group
}

func NewCacheGroup(name string, maxCacheByte int64, getter SourceGetter) *CacheGroup {
	if getter == nil {
		panic("sourceGetter is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &CacheGroup{
		name: name,
		mainCache: &cache{
			maxCacheBytes: maxCacheByte,
		},
		sourceGetter: getter,
		loader:       &singleflight.Group{},
	}
	groups[name] = g
	return g
}

func (g *CacheGroup) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(peer.Name(), key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

func (g *CacheGroup) load(key string) (value ByteView, err error) {
	val, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				value, err := g.getFromPeer(peer, key) // 远程获取
				if err == nil {
					log.Printf("[biuCache] success to get from peer [%s] [%s]", peer.Name(), key)
					return value, nil
				}
				log.Printf("[biuCache] failed to get from peer [%s] [%s] err [%s]", peer.Name(), key, err)
			} else {
				log.Printf("[biuCache] failed to get from peer [%s] no peer", key)
			}
		}
		return g.getLocally(key)
	})
	if err != nil {
		return val.(ByteView), nil
	}
	return
}

/**
1. 先读取缓存，如果缓存中不存在则跳转到2
2. 获取服务器列表，将请求分配到指定的服务器获取，如果指定服务器没有获取到，则跳转到3
3. 获取原始数据（从数据库读取或者文件之类的）
*/
func (g *CacheGroup) Get(key string) (view ByteView, err error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Printf("got [%s] from [%s] main cache", key, g.name)
		return v, nil
	}
	return g.load(key)
}

func (g *CacheGroup) GetName() string {
	return g.name
}

func (g *CacheGroup) getLocally(key string) (value ByteView, err error) {
	bytes, err := g.sourceGetter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value = ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return
}

func (g *CacheGroup) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

func (g *CacheGroup) RegisterPeers(peers PeerPicker) {
	g.peers = peers
}

func getGroup(name string) *CacheGroup {
	return groups[name]
}
