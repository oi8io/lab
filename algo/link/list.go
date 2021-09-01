package link

import (
	"errors"
	"fmt"
)

type Element int

type List struct {
	value Element
	Next  *List
}

func NewList(arr []int) *List {
	var list = new(List)
	var tmp *List
	for _, v := range arr {
		var i = &List{value: Element(v)}
		if tmp != nil {
			tmp.Next = i
			tmp = tmp.Next
		} else {
			list = i
			tmp = list
		}
	}
	return list
}

func (l *List) Print() {
	fmt.Println(l.value)
	if l.Next != nil {
		l.Next.Print()
	}
}

func (l *List) IsEmpty() bool {
	return l == nil
}

func (l *List) IsLast() bool {
	return l.Next == nil
}

func (l *List) Find(v Element) (*List, error) {
	if l.IsEmpty() {
		return nil, errors.New("not found")
	}
	if l.value == v {
		return l, nil
	} else {
		return l.Next.Find(v)
	}
}

func (l *List) FindPrevious(v Element) *List {
	x := l
	for {
		if !x.Next.IsEmpty() && x.Next.value != v {
			x = x.Next
		} else {
			break
		}
	}
	return x
}

func (l *List) Delete(v Element) {
	p := l.FindPrevious(v)
	if !p.Next.IsLast() {
		p.Next = p.Next.Next
	} else {
		p.Next = nil
	}
}

func (l *List) FindV1(v Element) *List {
	x := l
	for {
		if x.value != v && !x.IsEmpty() {
			x = x.Next
		} else {
			break
		}
	}
	return x
}

func (l *List) Insert(v Element, i *List) {
	new := &List{value: v}
	f, _ := l.Find(i.value)
	new.Next = f.Next
	f.Next = new
}
