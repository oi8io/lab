package link

import (
	"fmt"
)

type ArrayList interface {
	IsEmpty() bool
	IsLast(k int) bool
	IsFull() bool
	Expand()
	Get(k int) (interface{}, error)
	Set(k int, v interface{})
	Insert(k int, v interface{})
	Delete(k int)
	Print()
}

type Array struct {
	Size int
	list []interface{}
}

func NewArray() *Array {
	var list = new(Array)
	list.Size = 0
	list.list = make([]interface{}, 0, 10)
	return list
}

func (l *Array) Print() {
	fmt.Println(cap(l.list))
	fmt.Println(l.list)
}

func (l *Array) IsEmpty() bool {
	return l.Size == 0
}

func (l *Array) IsLast(k int) bool {
	return l.Size == k
}

func (l *Array) IsFull() bool {
	return l.Size == cap(l.list)
}

func (l *Array) Get(k int) (interface{}, error) {
	return l.list[k], nil
}

func (l *Array) Set(k int, v interface{}) {
	l.list[k] = v
}

func (l *Array) Expand() {
	i := cap(l.list)
	newList := make([]interface{}, 2*i)
	newList = append(newList, l.list...)
	l.list = newList
}

func (l *Array) Delete(k int) {
	l.list = append(l.list[:k-1], l.list[k+1:])
}

func (l *Array) BulkInsert(k int, v []interface{}) {
	l.Size = l.Size + len(v)
	if l.IsLast(k) {
		l.list = append(l.list, v...)
		return
	}
	l.list = append(l.list[:k], v...)
	l.list = append(l.list, l.list[k+1:]...)
}
func (l *Array) Insert(k int, v interface{}) {
	l.Size++
	if l.IsLast(k) {
		l.list = append(l.list, v)
		return
	}
	l.list = append(l.list[:k], v)
	l.list = append(l.list, l.list[k+1:]...)
}
