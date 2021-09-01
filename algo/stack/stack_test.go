package stack

import (
	"fmt"
	"testing"
)

var stack *Stack

func init() {
	stack = NewStack()
}

func TestStack_Pop(t *testing.T) {
	stack = stack.Push(5)
	fmt.Println(stack)

	stack = stack.Push(4)
	fmt.Println(stack)

	stack = stack.Push(2)
	fmt.Println(stack)

	stack = stack.Push(1)
	var x Element
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack, x = stack.Pop()
	fmt.Println(x)
	stack = stack.Push(9)
	stack = stack.Push(8)
	stack = stack.Push(7)
	stack = stack.Push(111)
	stack = stack.Push(118)
	stack, x = stack.Pop()
	fmt.Println(x)

}

func TestStack_Push(t *testing.T) {

}
