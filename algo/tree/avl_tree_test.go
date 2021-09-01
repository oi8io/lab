package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var avl *AVLTree

func init() {
	var arr  []int
	//arr := []int{9, 19, 8, 1, 3, 12, 18, 17, 22}
	//arr := []int{9, 10, 11, 12, 13, 14, 15, 16, 17}
	//arr := []int{9, 10, 11}

	//     8<-9<-10->11
	//arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1,12, 11,18,13,14,15,17,16}
	//arr := []int{10, 9, 8, 7, 6, 5, 8,12, 11,18,13,14,15,17,16}
	//arr := []int{12,11, 10, 9}
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//arr = []int{1, 2, 3, 4, 5, 6, 7,}
	arr = []int{1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13, 14, 15, 16,8,9}
	//arr := []int{8,    4,12,   2,6,10,14,    1,3,5,7,9,11,13,15}
	//for i := 1; i <= 10; i++ {
	//	arr= append(arr, i)
	//}
	avl = NewAVLTree(arr)

}
func TestAVLTree_For(t *testing.T) {

}

func TestAVLTree_For2(t *testing.T) {
	avl.Print()
}

func TestAVLTree_FindMax(t *testing.T) {
	node := avl.FindMax()
	fmt.Println(node.Value)
}

func TestAVLTree_FindMin(t *testing.T) {
	node := avl.FindMin()
	fmt.Println(node.Value)
}

func TestAVLTree_Find(t *testing.T) {
	//avl.For2(24)
	Convey("TestGetActionType", t, func() {
		node := avl.Find(25)
		So(node, ShouldBeNil)
	})
	Convey("TestGetActionType", t, func() {
		node := avl.Find(17)
		fmt.Println(node.Value)
		fmt.Println(node.Height)
		So(node.Value, ShouldBeGreaterThanOrEqualTo, 17)
		So(node.Height, ShouldBeGreaterThanOrEqualTo, 4)
		//So(node, ShouldBeNil)
	})
}
func TestAVLTree_Delete(t *testing.T) {
	Convey("Test Binary Search Tree", t, func() {
		//avl.Delete(9)
		avl.Print()
	})
}

func TestPrintTree(t *testing.T) {
	a := PrintTree(avl)
	for _, l := range a {
		fmt.Println(l)
	}
}
