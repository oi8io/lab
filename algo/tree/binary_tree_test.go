package tree

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	arr := []int{9, 19, 8, 1, 3, 12, 18, 17, 22}
	k := NewBinaryTree(arr)
	k.For2(12)
	//fmt.Println(x)
	//for i := 0; i < 24; i++ {
	//	for j := 0; j < 24; j++ {
	//		if x[i][j]==0 {
	//			fmt.Print(" "+ fmt.Sprintf("%3s"," "))
	//			continue
	//		}
	//		fmt.Print(" "+ fmt.Sprintf("%3d",x[i][j]))
	//	}
	//	fmt.Println()
	//}

}
