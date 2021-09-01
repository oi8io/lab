package base36

import (
	"fmt"
	"strconv"
)

//定义元素类型
type Element uint

//定义节点
type LinkNode struct {
	Data Element   //数据域
	Next *LinkNode //指针域，指向下一个节点
}

//函数接口
type LinkNoder interface {
	Add(head *LinkNode, new *LinkNode)              //后面添加
	Delete(head *LinkNode, index int)               //删除指定index位置元素
	Insert(head *LinkNode, index int, data Element) //在指定index位置插入元素
	GetLength(head *LinkNode) int                   //获取长度
	Search(head *LinkNode, data Element)            //查询元素的位置
	GetData(head *LinkNode, index int) Element      //获取指定index位置的元素
}

func Add(head *LinkNode, element Element) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	n := &LinkNode{element, nil}
	point.Next = n
}

func Delete(head *LinkNode, index int) () {

}
func Insert(head *LinkNode, index int, data Element) {

}

func GetLength(head *LinkNode) int {
	return 0
}

func GetData(head *LinkNode, index int) Element {
	return head.Data
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rs []int
	var res *ListNode
	var v1, v2 int
	carry := 0
	for {
		if l1 != nil {
			v1 = l1.Val
			fmt.Print(v1, ",")
			if l1.Next != nil {
				l1 = l1.Next
			} else {
				l1 = nil
			}
		} else {
			v1 = 0
		}

		if l2 != nil {
			v2 = l2.Val
			fmt.Print(v2, "][")
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2 = nil
			}
		} else {
			v2 = 0
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		rs = append(rs, sum%10)

		if l1 == nil && l2 == nil {
			if carry>0 {
				rs = append(rs, carry)
			}
			break
		}
	}
	for i:=len(rs)-1;i>=0;i-- {
		x := &ListNode{rs[i], nil}
		if res == nil {
			res = x
		} else {
			o := res
			x.Next = o
			res = x
		}
	}
	fmt.Println(rs)
	return res
}
func addTwoNumbersx(l1 *ListNode, l2 *ListNode) *ListNode {

	var num1 = 0
	var num2 = 0
	x := 1
	for {
		num1 += l1.Val * x
		x *= 10
		if l1.Next == nil {
			break
		} else {
			l1 = l1.Next
		}
	}

	x = 1
	for {
		num2 += l2.Val * x
		x *= 10
		if l2.Next == nil {
			break
		} else {
			l2 = l2.Next
		}
	}
	fmt.Println(num2)
	var r *ListNode
	sum := num1 + num2
	fmt.Println(sum)
	lenx := len(strconv.Itoa(sum))
	for i := lenx; i >= 0; i-- {
		xx := sum / pow(10, i)
		sum = sum % pow(10, i)
		fmt.Println(xx, sum)
		s := &ListNode{}
		s.Val = xx
		if r == nil {
			r = s
		} else {
			s.Next = r
			r = s
		}
	}
	return r
	//for {
	//	if sum > 0 {
	//		s := &ListNode{}
	//		s.Val = sum % 10
	//		sum = sum / 10
	//		if r == nil {
	//			r = s
	//		} else {
	//			s.Next = r
	//			r = s
	//		}
	//	} else {
	//		break
	//	}
	//}
	//return r
	//fmt.Println(r)
	x = 1
	var num3 = 0
	for {
		num3 += r.Val * x
		x *= 10
		if r.Next == nil {
			break
		} else {
			r = r.Next
		}
	}
	fmt.Println(num3)
	return r

}
