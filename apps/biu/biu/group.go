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
	name      string
	mainCache iCache
	getter    Getter
	peers     PeerPicker
	loader    *singleflight.Group
}

func NewCacheGroup(name string, maxCacheByte int64, getter Getter) *CacheGroup {
	if getter == nil {
		panic("getter is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &CacheGroup{
		name: name,
		mainCache: &cache{
			maxCacheBytes: maxCacheByte,
		},
		getter: getter,
		loader: &singleflight.Group{},
	}
	groups[name] = g
	return g
}

func (g *CacheGroup) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

func (g *CacheGroup) load(key string) (value ByteView, err error) {
	val, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err := g.getFromPeer(peer, key); err == nil {
					log.Printf("[biuCache] success to get from peer [%s] [%s]", peer.Name(), key)
					return value, nil
				}
				log.Println("[biuCache] Failed to get from peer", err)
			}
		}
		return g.getLocally(key)
	})
	if err != nil {
		return val.(ByteView), nil
	}
	return
}

func (g *CacheGroup) Get(key string) (view ByteView, err error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	return g.load(key)
}

func (g *CacheGroup) getLocally(key string) (value ByteView, err error) {
	bytes, err := g.getter.Get(key)
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
