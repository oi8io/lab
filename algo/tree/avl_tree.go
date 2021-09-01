package tree

import (
	"fmt"
	"strconv"
	"strings"
	//"golang.org/x/text/encoding/charmap"
)

const arrLen = 4096 * 2

type CommonTree interface {
	GetHeight() int
	SetHeight()
	GetValue(v int)
	SetValue() int
	GetDepth() int
	GetRight() *CommonTree
	GetLeft() *CommonTree
	Print()
}

var Avltree_SearchArray [arrLen][arrLen]string

type AVLTreeElement struct {
	Value   int
	Height  int
	Depth   int
	PayLoad int32
}

// todo use interface
type AVLTree struct {
	AVLTreeElement
	Left  *AVLTree
	Right *AVLTree
}

func (avl *AVLTree) SetHeight() {
	panic("implement me")
}

func (avl *AVLTree) GetValue(v int) {
	panic("implement me")
}

func (avl *AVLTree) SetValue() int {
	panic("implement me")
}

func (avl *AVLTree) GetDepth() int {
	panic("implement me")
}

func (avl *AVLTree) GetRight() *CommonTree {
	panic("implement me")
}

func (avl *AVLTree) GetLeft() *CommonTree {
	panic("implement me")
}

func NewAVLTree(values []int) *AVLTree {
	var t *AVLTree
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
func (bt *AVLTree) Insert(v int) *AVLTree {
	tmp := new(AVLTree)
	tmp.Value = v
	if bt == nil {
		bt = tmp
	} else {
		if v == bt.Value {
			return bt
		} else if v > bt.Value {
			bt.Right = bt.Right.Insert(v)
			//fmt.Println("insert", bt.right.value)
			//右边比左边高
			if bt.Right.GetHeight()-bt.Left.GetHeight() == 2 {
				//fmt.Println("right more ", bt.value)
				if v > bt.Right.Value { // 右右
					//fmt.Println("right ", bt.value)
					bt = bt.SingleRotateWithLeft()
				} else { // 右左
					//fmt.Println("right left ", bt.value)
					bt = bt.DoubleRotateWithLeft()
				}
			}
		} else {
			bt.Left = bt.Left.Insert(v)
			if bt.Left.GetHeight()-bt.Right.GetHeight() == 2 {
				if v < bt.Left.Value { // 左左

					bt.Print()
					fmt.Println("SingleRotateWithRight")
					//fmt.Println("left ", bt.value)
					bt = bt.SingleRotateWithRight()
					bt.Print()
				} else { //左右
					//fmt.Println("left right ", bt.value)
					bt.Print()
					fmt.Println("DoubleRotateWithRight")
					bt = bt.DoubleRotateWithRight()
					bt.Print()
				}
			}
		}
	}
	// 设置高度
	bt.GetHeightByChild()
	return bt
}
func (t *AVLTree) SingleRotateWithLeft() *AVLTree {
	tmp := t.Right
	t.Right = tmp.Left
	t.GetHeightByChild()
	tmp.Left = t
	tmp.GetHeightByChild()
	return tmp
}

func (bt *AVLTree) SingleRotateWithRight() *AVLTree {
	tmp := bt.Left
	bt.Left = tmp.Right
	bt.GetHeightByChild()
	tmp.Right = bt
	tmp.GetHeightByChild()
	return tmp
}

//DoubleRotateWithLeft 右子树的左边高，先向右旋转，再向左旋转
func (bt *AVLTree) DoubleRotateWithLeft() *AVLTree { //左边太深了
	bt.Right = bt.Right.SingleRotateWithRight()
	return bt.SingleRotateWithLeft()
}
func (bt *AVLTree) DoubleRotateWithRight() *AVLTree { //左边太深了
	bt.Left = bt.Left.SingleRotateWithLeft()
	return bt.SingleRotateWithRight()
}

func (bt *AVLTree) GetHeight() int {
	if bt == nil {
		return 0
	}
	return bt.Height
}

func (bt *AVLTree) GetHeightByChild() int {
	bt.Height = max(bt.Left.GetHeight(), bt.Right.GetHeight()) + 1
	return bt.Height
}

func (bt *AVLTree) Delete(v int) *AVLTree {
	if bt == nil {
		return nil
	}
	if bt.Value == v {
		if bt.Left == nil {
			if bt.Right == nil {
				return nil
			} else {
				return bt.Right
			}
		} else {
			if bt.Right == nil {
				return bt.Left
			} else {
				min := bt.Right.FindMin()
				bt.Value = min.Value
				bt.Right = bt.Right.Delete(min.Value)
				return bt
			}
		}
	}
	if v > bt.Value {
		bt.Right = bt.Right.Delete(v)
		return bt
	} else {
		bt.Left = bt.Left.Delete(v)
		return bt
	}
}

func (b *AVLTree) For(cur int) {
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

func (b *AVLTree) Show(m map[int32]string, pre string) map[int32]string {
	if b == nil {
		return m
	}
	if b.PayLoad != 0 {
		m[b.PayLoad] = pre
		//fmt.Printf("%#b\n", pre)
		//fmt.Println(b.PayLoad, "=>", pre)
	}
	//pre := strings.Repeat("  ", b.Height) + "|->"
	if b.Left != nil {
		//fmt.Println(pre+"<-|")
		b.Left.Show(m, pre+"0")
	}
	//fmt.Println(pre + strconv.Itoa(b.Value))
	if b.Right != nil {
		//fmt.Println(pre+"|->")
		b.Right.Show(m, pre+"1")
	}
	return m
}

func (b *AVLTree) FindMax() *AVLTree {
	if b.Right != nil {
		return b.Right.FindMax()
	} else {
		return b
	}
}

func (b *AVLTree) FindMin() *AVLTree {
	if b.Left != nil {
		return b.Left.FindMin()
	} else {
		return b
	}
}
func (b *AVLTree) Find(v int) *AVLTree {
	if b.Value == v {
		return b
	}
	if b.Value > v {
		if b.Left == nil {
			return nil
		}
		return b.Left.Find(v)
	} else {
		if b.Right == nil {
			return nil
		}
		return b.Right.Find(v)
	}
}

const (
	leftChar  = " ┌─"
	rightChar = "─┐ "
	midChar   = "───"
	eptChar   = "   "
)

// ┐

//  ⌐ ¬ https://en.wikipedia.org/wiki/Code_page_437
func (b *AVLTree) Print() {
	cur := pow(2, b.Height)
	b.For2(b.Height, cur)
	arr := Avltree_SearchArray
	Avltree_SearchArray = [arrLen][arrLen]string{}

	for i := b.Height; i >= 0; i-- {
		for j := 0; j <= pow(2, b.Height+1); j++ {
			val := arr[i][j]
			if val == "" {
				fmt.Print(eptChar)
			} else {
				fmt.Print(val)
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

func (b *AVLTree) For2(height, cur int) {
	if b == nil {
		return
	}
	mov := pow(2, b.Height-1)
	val := strconv.Itoa(b.Value + 100)
	Avltree_SearchArray[height][cur] = val
	if height > arrLen || (cur+mov) > arrLen || cur < mov {
		//fmt.Println(height, cur, mov)
		return
	}
	//pre := strings.Repeat("  ", b.height) + "|->"
	if b.Left != nil {
		//fmt.Println(pre+"<-|")

		Avltree_SearchArray[height][cur-mov] = leftChar
		for i := cur - mov + 1; i < cur; i++ {
			Avltree_SearchArray[height][i] = midChar
		}
		b.Left.For2(height-1, cur-mov) //4
	}
	//fmt.Println(pre + strconv.Itoa(b.value))
	if b.Right != nil {
		//fmt.Println(pre+"|->")
		for i := cur + 1; i < cur+mov; i++ {
			Avltree_SearchArray[height][i] = midChar
		}
		Avltree_SearchArray[height][cur+mov] = rightChar
		b.Right.For2(height-1, cur+mov) //12
	}
}
