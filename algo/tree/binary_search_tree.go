package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type BinarySearchTreeElement struct {
	value  int
	height int
}

type BinarySearchTree struct {
	BinarySearchTreeElement
	left  *BinarySearchTree
	right *BinarySearchTree
}

func (bt *BinarySearchTree) Insert(v int) {
	tmp := new(BinarySearchTree)
	tmp.value = v
	if bt == nil {
		tmp.height = 0
		bt = tmp
		return
	}
	tmp.height = bt.height + 1
	if v > bt.value {
		if bt.right != nil {
			bt.right.Insert(v)
		} else {
			bt.right = tmp
		}
	} else {
		if bt.left != nil {
			bt.left.Insert(v)
		} else {
			bt.left = tmp
		}
	}
}

func (bt *BinarySearchTree) Delete(v int) *BinarySearchTree {
	if bt == nil {
		return nil
	}
	if bt.value == v {
		if bt.left == nil {
			if bt.right == nil {
				return nil
			} else {
				return bt.right
			}
		} else {
			if bt.right == nil {
				return bt.left
			} else {
				min := bt.right.FindMin()
				bt.value = min.value
				bt.right = bt.right.Delete(min.value)
				return bt
			}
		}
	}
	if v > bt.value {
		bt.right = bt.right.Delete(v)
		return bt
	} else {
		bt.left = bt.left.Delete(v)
		return bt
	}
}

func NewBinarySearchTree(values []int) *BinarySearchTree {
	var t *BinarySearchTree
	for _, v := range values {
		//fmt.Println("insert v", v)
		if t == nil {
			t = new(BinarySearchTree)
			t.value = v
		} else {
			t.Insert(v)
		}
	}
	return t
}

func (b *BinarySearchTree) For(cur int) {
	if b == nil {
		return
	}
	pre := strings.Repeat("  ", b.height) + "|->"
	if b.left != nil {
		//fmt.Println(pre+"<-|")
		b.left.For(cur - 1)
	}
	fmt.Println(pre + strconv.Itoa(b.value))
	if b.right != nil {
		//fmt.Println(pre+"|->")
		b.right.For(cur - 1)
	}
}

func (b *BinarySearchTree) FindMax() *BinarySearchTree {
	if b.right != nil {
		return b.right.FindMax()
	} else {
		return b
	}
}

func (b *BinarySearchTree) FindMin() *BinarySearchTree {
	if b.left != nil {
		return b.left.FindMin()
	} else {
		return b
	}
}
func (b *BinarySearchTree) Find(v int) *BinarySearchTree {
	if b.value == v {
		return b
	}
	if b.value > v {
		if b.left == nil {
			return nil
		}
		return b.left.Find(v)
	} else {
		if b.right == nil {
			return nil
		}
		return b.right.Find(v)
	}

}

func (b *BinarySearchTree) Print() {
	b.For2(12)
	for i := 0; i < 24; i++ {
		for j := 0; j < 24; j++ {
			if searchArray[i][j] == 0 {
				fmt.Print(" " + fmt.Sprintf("%3s", " "))
				continue
			}
			fmt.Print(" " + fmt.Sprintf("%3d", searchArray[i][j]))
		}
		fmt.Println()
	}
}

var searchArray = [24][24]int{}
var maxHeight int
func (b *BinarySearchTree) For2(cur int) {
	if b == nil {
		return
	}
	//pre := strings.Repeat("  ", b.height) + "|->"
	if b.left != nil {
		//fmt.Println(pre+"<-|")
		b.left.For2(cur - 2)
	}
	searchArray[b.height][cur] = b.value+100
	//fmt.Println(pre + strconv.Itoa(b.value))
	if b.right != nil {
		//fmt.Println(pre+"|->")
		b.right.For2(cur + 2)
	}
}
