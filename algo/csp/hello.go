package csp

import (
	"fmt"
	"time"
)

// say hello every 1s
func SayHello()  {
	for  {
		fmt.Println("Hello Golang")
		time.Sleep(time.Second)
	}
}
// say hello every 1s
func SayHi()  {
	for  {
		fmt.Println("Hi Golang")
		time.Sleep(time.Second)
	}
}

