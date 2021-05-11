package biu

import (
	"oi.io/apps/biu/biu/lru"
	"sync"
)

type iCache interface {
	getLru() lru.CacheLru
	add(key string, view ByteView)
	get(key string) (view ByteView, ok bool)
	del(key string)
}

type cache struct {
	lru           lru.CacheLru
	mu            sync.Mutex
	maxCacheBytes int64
}


func (c *cache) getLru() lru.CacheLru {
	if c.lru == nil {
		c.lru = lru.NewCache(c.maxCacheBytes, nil)
	}
	return c.lru
}

func (c *cache) add(key string, view ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.getLru().Add(key, view)
}

func (c *cache) get(key string) (view ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, ok := c.getLru().Get(key); ok {
		return value.(ByteView), ok
	}
	return
}

func (c *cache) del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.getLru().Del(key)
}
