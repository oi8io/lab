package biu

// 原始数据读取方法
type SourceGetter interface {
	Get(key string) (value []byte, err error)
}

// A SourceGetter loads data for a key.
type GetterFunc func(key string) (value []byte, err error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
