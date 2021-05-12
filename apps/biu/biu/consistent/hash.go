package consistent

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

// 一致性hash

type hashFunc func(b []byte) uint32

//
//定义了函数类型 Hash，采取依赖注入的方式，允许用于替换成自定义的 Hash 函数，也方便测试时替换，默认为 crc32.ChecksumIEEE 算法。
//Map 是一致性哈希算法的主数据结构，包含 4 个成员变量：Hash 函数 hash；虚拟节点倍数 replicas；哈希环 keys；虚拟节点与真实节点的映射表 hashMap，键是虚拟节点的哈希值，值是真实节点的名称。
type HashMap struct {
	hashFunc hashFunc       // 默认为 crc32.ChecksumIEEE 算法
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环
	hashMap  map[int]string //虚拟节点与真实节点的映射表
}

//构造函数 NewHashMap() 允许自定义虚拟节点倍数和 Hash 函数。
func NewHashMap(replicas int, hashFunc hashFunc, ) *HashMap {
	if hashFunc == nil {
		hashFunc = crc32.ChecksumIEEE
	}
	return &HashMap{hashFunc: hashFunc, replicas: replicas, hashMap: make(map[int]string)}
}

// Add adds some keys to the hash.
func (m *HashMap) Add(nodeMap map[string]string) {
	for key, _ := range nodeMap {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Add adds some keys to the hash.
func (m *HashMap) Get(key string) (nodeName string, err error) {
	if key == "" {
		return "", fmt.Errorf("key is required")
	}
	n := int(m.hashFunc([]byte(key)))
	hit := sort.Search(len(m.keys), func(i int) bool { //   传入长度 循环搜索
		return m.keys[i] >= n
	})
	idx := m.keys[hit%len(m.keys)]
	return m.hashMap[idx], nil
}
