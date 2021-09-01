package main

import "fmt"

func main() {
	type Test struct{}
	v := Test{}
	Print(v)
}

func Print(v interface{}) {
	println(v)
	fmt.Printf("%+v", v)
}
