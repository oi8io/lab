package tree

import (
	"fmt"
	"strconv"
	"strings"
	//"golang.org/x/text/encoding/charmap"
)

var SplayTree_SearchArray = [20][512]string{}

type SplayTreeElement struct {
	value  int
	height int
	depth  int
}

// todo use interface
type SplayTree struct {
	SplayTreeElement
	left  *SplayTree
	right *SplayTree
}

func NewSplayTree(values []int) *SplayTree {
	var t *SplayTree
	for _, v := range values {
		t = t.Insert(v)
	}
	return t
}

/**
1. 左左插入
2. 左右插入
3. 右左插入
4. 右右插入
*/
func (bt *SplayTree) Insert(v int) *SplayTree {
	tmp := new(SplayTree)
	tmp.value = v
	if bt == nil {
		bt = tmp
	} else {
		if v == bt.value {
			return bt
		} else if v > bt.value {
			bt.right = bt.right.Insert(v)
			//fmt.Println("insert", bt.right.value)
			//右边比左边高
			if bt.right.Height()-bt.left.Height() == 2 {
				//fmt.Println("right more ", bt.value)
				if v > bt.right.value { // 右右
					//fmt.Println("right ", bt.value)
					bt = bt.SingleRotateWithLeft()
				} else { // 右左
					//fmt.Println("right left ", bt.value)
					bt = bt.DoubleRotateWithLeft()
				}
			}
		} else {
			bt.left = bt.left.Insert(v)
			if bt.left.Height()-bt.right.Height() == 2 {
				if v < bt.left.value { // 左左
					//fmt.Println("left ", bt.value)
					bt = bt.SingleRotateWithRight()
				} else { //左右
					//fmt.Println("left right ", bt.value)
					bt = bt.DoubleRotateWithRight()
				}
			}
		}
	}
	// 设置高度
	bt.SetHeight()
	return bt
}

func (bt *SplayTree) SingleRotateWithLeft() *SplayTree {
	tmp := bt.right
	bt.right = tmp.left
	bt.SetHeight()
	tmp.left = bt
	tmp.SetHeight()
	return tmp
}

func (bt *SplayTree) SingleRotateWithRight() *SplayTree {
	tmp := bt.left
	bt.left = tmp.right
	bt.SetHeight()
	tmp.right = bt
	tmp.SetHeight()
	return tmp
}

//DoubleRotateWithLeft 右子树的左边高，先向右旋转，再向左旋转
func (bt *SplayTree) DoubleRotateWithLeft() *SplayTree { //左边太深了
	bt.right = bt.right.SingleRotateWithRight()
	return bt.SingleRotateWithLeft()
}
func (bt *SplayTree) DoubleRotateWithRight() *SplayTree { //左边太深了
	bt.left = bt.left.SingleRotateWithLeft()
	return bt.SingleRotateWithRight()
}

func (bt *SplayTree) Height() int {
	if bt == nil {
		return 0
	}
	return bt.height
}

func (bt *SplayTree) SetHeight() int {
	bt.height = max(bt.left.Height(), bt.right.Height()) + 1
	return bt.height
}

func (bt *SplayTree) Delete(v int) *SplayTree {
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

func (b *SplayTree) For(cur int) {
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

func (b *SplayTree) FindMax() *SplayTree {
	if b.right != nil {
		return b.right.FindMax()
	} else {
		return b
	}
}

func (b *SplayTree) FindMin() *SplayTree {
	if b.left != nil {
		return b.left.FindMin()
	} else {
		return b
	}
}
func (b *SplayTree) Find(v int) *SplayTree {
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

// ┐

//  ⌐ ¬ https://en.wikipedia.org/wiki/Code_page_437
func (b *SplayTree) Print() {
	cur := pow(2, b.height)
	b.For2(cur)
	arr := SplayTree_SearchArray
	SplayTree_SearchArray = [20][512]string{}

	for i := b.height; i >= 0; i-- {
		sub := false
		for j := 0; j <= pow(2, b.height+1); j++ {
			val := arr[i][j]
			if val == "" {
				if sub {
					fmt.Print(midChar)
				} else {
					fmt.Print("   ")
				}
			} else {
				fmt.Print(val)
				if val == leftChar {
					sub = true
				} else if val == rightChar {
					sub = false
				}
			}
			//
			//if val == 0 {
			//	fmt.Print(fmt.Sprintf(padding, "*"))
			//	//continue
			//} else {
			//	fmt.Println(val)
			//	fmt.Print(fmt.Sprintf(padding, "-"))
			//}
		}
		fmt.Println()
	}
}

func (b *SplayTree) For2(cur int) {
	if b == nil {
		return
	}
	mov := pow(2, b.height-1)
	val := strconv.Itoa(b.value + 100)
	SplayTree_SearchArray[b.height][cur] = val
	//pre := strings.Repeat("  ", b.height) + "|->"
	if b.left != nil {
		//fmt.Println(pre+"<-|")
		SplayTree_SearchArray[b.height][cur-mov] = leftChar
		b.left.For2(cur - mov) //4
	}
	//fmt.Println(pre + strconv.Itoa(b.value))
	if b.right != nil {
		//fmt.Println(pre+"|->")
		SplayTree_SearchArray[b.height][cur+mov] = rightChar
		b.right.For2(cur + mov) //12
	}
}
