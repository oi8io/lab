package main

import (
	"fmt"
	"os"
	"time"
)

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray)

	aArray[0] = 10
	return
}

func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func main_def() {
	defer fmt.Println("defer main") // will this be printed when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if i := recover(); i != nil {
				fmt.Println("sfsafsadfasdfsa",i)
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
		}()
	}()

	time.Sleep(11 * time.Second)
	//fmt.Printf("get result %d\r\n", result)
}
