package cache

import "time"

//说明
//问题 - ⼀一个简易易的内存缓存系统 该程序需要满⾜足以下要求:
//1. 支持设定过期时间，精度为秒级。
//2. 支持设定最⼤大内存，当内存超出时候做出合理理的处理理。
//3. 支持并发安全。
//4. 为简化编程细节，⽆无需实现数据落地。
/**
支持过期时间和最⼤内存⼤小的的内存缓存库。
*/
type Cache interface {
	//size 是一个字符串串。⽀支持以下参数: 1KB，100KB，1MB，2MB，1GB 等
	SetMaxMemory(size string) bool
	// 设置一个缓存项，并且在expire时间之后过期
	Set(key string, val interface{}, expire time.Duration)
	// 获取一个值
	Get(key string) (interface{}, bool)
	// 删除一个值
	Del(key string) bool
	// 检测一个值 是否存在
	Exists(key string) bool
	// 清空所有值
	Flush() bool
	// 返回所有的key数量
	Keys() int64
}


/**
存储系统：
1. 固定可以设置大小
2. 当大小不足以进行当前空间时执行回收
3. 当查询到过期时进行回收

 */
type Store struct {

}

type HyperCache struct {
}

func (h HyperCache) Set(key string, val interface{}, expire time.Duration) {
	panic("implement me")
}

func (h HyperCache) Get(key string) (interface{}, bool) {
	panic("implement me")
}

func (h HyperCache) Del(key string) bool {
	panic("implement me")
}

func (h HyperCache) Exists(key string) bool {
	panic("implement me")
}

func (h HyperCache) Flush() bool {
	panic("implement me")
}

func (h HyperCache) Keys() int64 {
	panic("implement me")
}


