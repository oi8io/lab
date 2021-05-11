package biu

type iByteView interface {
	Len() int64
	ByteSlice() []byte
	String() string
}

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int64 {
	return int64(len(v.b))
}

func (v *ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v *ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	r := make([]byte, len(b))
	copy(r, b)
	return r
}
