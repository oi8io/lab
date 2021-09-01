package link

import (
	"fmt"
	"testing"
)

var list *List

func init() {
	arr := []int{1, 2, 4, 8, 12, 18}
	list = NewList(arr)
}
func TestNewList(t *testing.T) {
	list.Print()
}

func TestList_Find(t *testing.T) {
	i := list.FindV1(4)
	fmt.Println(i)
}

func TestList_Insert(t *testing.T) {
	i, _ := list.Find(4)
	list.Insert(199, i)
	list.Print()
}

func TestList_FindPrevious(t *testing.T) {
	i := list.FindPrevious(4)
	i.Print()
}

func TestList_Delete(t *testing.T) {
	list.Delete(18)
	list.Print()
}
