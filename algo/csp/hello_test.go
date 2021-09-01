package csp

import (
	"fmt"
	"testing"
	"time"
)

func TestSayHello(t *testing.T) {
	go SayHello()
	SayHi()
	time.Sleep(time.Hour*10)
}

func TestSayHi(t *testing.T) {
	go SayHello()
	go SayHi()
	time.Sleep(time.Hour*10)
}

func TestSayHello2(t *testing.T) {
	go func() {
		for  {
			fmt.Println("Hello Golang")
			time.Sleep(time.Second)
		}
	}()
	func() {
		for  {
			fmt.Println("Hi Golang")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Hour*10)
}
