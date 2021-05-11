package lru

import (
	"container/list"
	"sync"
)

type CacheLru interface {
	Get(key string) (value Value, ok bool)
	Add(key string, value Value)
	Del(key string)
	Evict()
	Len() int64
}

//CacheLruBase cache 结构
type CacheLruBase struct {
	maxBytes  int64                    // 最大字节数
	usedBytes int64                    // 已经使用字节数
	list      *list.List               // 双向链表（builtin）
	cache     map[string]*list.Element // 字典定义，键是字符串，值是双向链表中对应节点的指针
	mu        sync.Mutex
	OnEvicted func(key string, value Value) // 是某条记录被移除时的回调函数，可以为 nil。
}

func NewCache(maxBytes int64, onEvicted func(key string, value Value)) *CacheLruBase {
	return &CacheLruBase{maxBytes: maxBytes, list: list.New(), cache: make(map[string]*list.Element), OnEvicted: onEvicted}
}

func (c *CacheLruBase) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, ok
	}
	return
}

func (c *CacheLruBase) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.usedBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.list.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.usedBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes > 0 && c.maxBytes < c.usedBytes {
		c.Evict() // todo 先淘汰再插入
	}
}

func (c *CacheLruBase) Del(key string) {
	if ele, ok := c.cache[key]; ok {
		c.list.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.usedBytes += int64(kv.value.Len())
	}
}

func (c *CacheLruBase) Evict() {
	back := c.list.Back()
	if back == nil {
		return
	}
	c.list.Remove(back)
	kv := back.Value.(*entry)
	delete(c.cache, kv.key)
	c.usedBytes -= int64(len(kv.key)) + kv.value.Len()
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *CacheLruBase) Len() int64 {
	return int64(c.list.Len())
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int64
}
