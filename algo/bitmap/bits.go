package bitmap

import "fmt"

type BitMap struct {
	bits  map[uint]uint64
}

func NewBitMap() *BitMap {
	return &BitMap{make(map[uint]uint64)}
}

func (bitmap *BitMap) Add(num uint) {
	index := num >> 7
	pos := num % 64
	fmt.Println(pos)
	bitmap.bits[index] |= 1 << pos
}
func (bitmap *BitMap) Has(num uint) bool {
	index := num >> 7
	pos := num % 64
	exist := (bitmap.bits[index]&(1 << pos))>>pos
	if exist == 1 {
		return true
	}
	return  false
}

//ByteToBinaryString函数来源:
// Go语言版byte变量的二进制字符串表示
// http://www.sharejs.com/codes/go/4357
func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1

		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}

		data <<= 1
	}
	return str
}
