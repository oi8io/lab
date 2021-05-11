package biu

import (
	"fmt"
	"oi.io/apps/biu/biu/lru"
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
	}
	groups[name] = g
	return g
}

func (g *CacheGroup) getLru() lru.CacheLru {
	panic("implement me")
}

func (g *CacheGroup) load(key string) (view ByteView, err error) {
	return g.getLocally(key)
}

func (g *CacheGroup) Get(key string) (view ByteView, err error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	if v, ok := g.mainCache.get(key); ok {
		fmt.Println("cache hit", key)
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
