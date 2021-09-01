package tree

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var arr []int
var k *BinarySearchTree

func init() {
	var arr  []int
	for i := 1; i <= 12; i++ {
		arr= append(arr, i)
	}
	k = NewBinarySearchTree(arr)
}


func TestBinarySearchTree_For2(t *testing.T) {

}

func TestBinarySearchTree_FindMax(t *testing.T) {
	node := k.FindMax()
	fmt.Println(node.value)
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	node := k.FindMin()
	fmt.Println(node.value)
}

func TestBinarySearchTree_Find(t *testing.T) {
	//k.For2(24)
	Convey("TestGetActionType", t, func() {
		node := k.Find(25)
		So(node, ShouldBeNil)
	})
	Convey("TestGetActionType", t, func() {
		node := k.Find(17)
		fmt.Println(node.value)
		fmt.Println(node.height)
		So(node.value, ShouldBeGreaterThanOrEqualTo, 17)
		So(node.height, ShouldBeGreaterThanOrEqualTo, 4)
		//So(node, ShouldBeNil)
	})
}
func TestBinarySearchTree_Delete(t *testing.T) {
	Convey("Test Binary Search Tree", t, func() {
		//k.Delete(9)
		k.Print()
	})
}
