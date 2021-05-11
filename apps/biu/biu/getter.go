package biu

type Getter interface {
	Get(key string) (value []byte, err error)
}
// A Getter loads data for a key.
type GetterFunc func(key string) (value []byte, err error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}


