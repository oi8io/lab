package main

import (
	"fmt"
	"time"
)

func A() {
	fmt.Println("我是A,该来的还是来了")
	panic("我是A，现在我撕票了")
}
func B() {
	fmt.Println("我是B,该来的还是来了")
	panic("我是B，现在我撕票了")
}

func C() {
	if i := recover(); i != nil {
		fmt.Println("panic:", i)
		fmt.Println("recover: 我是C，我是来救场的")
	}
	time.Sleep(time.Second * 10)
}

func main() {
	defer A()
	defer B()
	defer C()
	panic("我是主，现在我撕票了")
	fmt.Println("撕票后")
}
