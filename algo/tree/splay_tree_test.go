package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var spt *SplayTree

func init() {
	//arr := []int{9, 19, 8, 1, 3, 12, 18, 17, 22}
	//arr := []int{9, 10, 11, 12, 13, 14, 15, 16, 17}
	//arr := []int{9, 10, 11}

	//     8<-9<-10->11
	//arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 12, 11, 18, 13, 14, 15, 17, 16}
	//arr := []int{10, 9, 8, 7, 6, 5, 8,12, 11,18,13,14,15,17,16}
	//arr := []int{12,11, 10, 9}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 8, 9, 13, 14, 15, 16,}
	//arr := []int{8,    4,12,   2,6,10,14,    1,3,5,7,9,11,13,15}
	//spt = NewSplayTree(arr)

}
func TestSplayTree_For(t *testing.T) {

}

func TestSplayTree_For2(t *testing.T) {
	spt.Print()
}

func TestSplayTree_FindMax(t *testing.T) {
	node := spt.FindMax()
	fmt.Println(node.value)
}

func TestSplayTree_FindMin(t *testing.T) {
	node := spt.FindMin()
	fmt.Println(node.value)
}

func TestSplayTree_Find(t *testing.T) {
	//spt.For2(24)
	Convey("TestGetActionType", t, func() {
		node := spt.Find(25)
		So(node, ShouldBeNil)
	})
	Convey("TestGetActionType", t, func() {
		node := spt.Find(17)
		fmt.Println(node.value)
		fmt.Println(node.height)
		So(node.value, ShouldBeGreaterThanOrEqualTo, 17)
		So(node.height, ShouldBeGreaterThanOrEqualTo, 4)
		//So(node, ShouldBeNil)
	})
}
func TestSplayTree_Delete(t *testing.T) {
	Convey("Test Binary Search Tree", t, func() {
		spt.Delete(9)
		spt.Print()
	})
}
