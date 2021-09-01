package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type BinaryTree struct {
	Value  int
	Height int
	Left   *BinaryTree
	Right  *BinaryTree
}

func (t *BinaryTree)PreOrderTraversal () {
	fmt.Println(strconv.Itoa(t.Value))
	if t.Left != nil {
		t.Left.PreOrderTraversal()
	}
	if t.Right != nil {
		t.Right.PreOrderTraversal()
	}
}

func (t *BinaryTree)InOrderTraversal () {
	if t.Left != nil {
		t.Left.PreOrderTraversal()
	}
	fmt.Println(strconv.Itoa(t.Value))
	if t.Right != nil {
		t.Right.PreOrderTraversal()
	}
}
func (t *BinaryTree)PostOrderTraversal () {
	if t.Left != nil {
		t.Left.PreOrderTraversal()
	}
	if t.Right != nil {
		t.Right.PreOrderTraversal()
	}
	fmt.Println(strconv.Itoa(t.Value))
}

func (bt *BinaryTree) Insert(v int) {
	tmp := new(BinaryTree)
	tmp.Value = v
	if bt == nil {
		tmp.Height = 0
		bt = tmp
		return
	}
	tmp.Height = bt.Height + 1
	if v > bt.Value {
		if bt.Right != nil {
			bt.Right.Insert(v)
		} else {
			bt.Right = tmp
		}
	} else {
		if bt.Left != nil {
			bt.Left.Insert(v)
		} else {
			bt.Left = tmp
		}
	}
}

func NewBinaryTree(values []int) *BinaryTree {
	var t *BinaryTree
	for _, v := range values {
		fmt.Println("insert v", v)
		if t == nil {
			t = new(BinaryTree)
			t.Value = v
		} else {
			t.Insert(v)
		}
	}
	return t
}

func (b *BinaryTree) For(cur int) {
	if b == nil {
		return
	}
	pre := strings.Repeat("  ", b.Height) + "|->"
	if b.Left != nil {
		//fmt.Println(pre+"<-|")
		b.Left.For(cur - 1)
	}
	fmt.Println(pre + strconv.Itoa(b.Value))
	if b.Right != nil {
		//fmt.Println(pre+"|->")
		b.Right.For(cur - 1)
	}

}

var x = [24][24]int{}

func (b *BinaryTree) For2(cur int) {
	if b == nil {
		return
	}

	//pre := strings.Repeat("  ", b.height) + "|->"
	if b.Left != nil {
		//fmt.Println(pre+"<-|")
		b.Left.For2(cur - 2)
	}
	x[b.Height][cur] = b.Value
	//fmt.Println(pre + strconv.Itoa(b.value))
	if b.Right != nil {
		//fmt.Println(pre+"|->")
		b.Right.For2(cur + 2)
	}
}

func (b *BinaryTree) For3(cur int) {
	if b == nil {
		return
	}

	//pre := strings.Repeat("  ", b.height) + "|->"
	if b.Left != nil {
		//fmt.Println(pre+"<-|")
		b.Left.For2(cur - 2)
	}
	x[b.Height][cur] = b.Value
	//fmt.Println(pre + strconv.Itoa(b.value))
	if b.Right != nil {
		//fmt.Println(pre+"|->")
		b.Right.For2(cur + 2)
	}
}


