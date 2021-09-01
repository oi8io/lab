package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{19, 28, 17, 68, 45, 16, 44, 48}
	//arr := []int{5, 2, 4}
	//InsertionSort(arr)
	InsertSort(arr)
	fmt.Println(arr)
}

func TestFb(t *testing.T) {
	Fb(50)
}

func TestSelectSort(t *testing.T) {
	arr := []int{19, 28, 17, 68, 45, 16, 44, 48}
	a := SelectSort(arr)
	fmt.Println(a)
	arr = []int{19}
	a = SelectSort(arr)
	fmt.Println(a)
}
