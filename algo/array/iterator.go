package link

import (
	"errors"
)

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}
type Iterable interface {
	Iterator() Iterator
}
type ArrayListIterator struct {
	list         *Array
	currentIndex int
}

func (l *Array) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.list = l
	it.currentIndex = 0
	return it
}

func (a *ArrayListIterator) HasNext() bool {
	return a.list.Size > a.currentIndex
}

func (a *ArrayListIterator) Next() (interface{}, error) {
	if !a.HasNext() {
		return nil, errors.New("no next")
	}
	idx := a.currentIndex
	a.currentIndex++
	return a.list.Get(idx)
}

func (a *ArrayListIterator) Remove() {
	a.currentIndex--
	a.list.Delete(a.currentIndex)
}

func (a *ArrayListIterator) GetIndex() int {
	panic("implement me")
}
