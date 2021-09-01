package main

import (
	"fmt"
	"strconv"
)

func mainx()  {
	hash := make(map[string]int, 26)
	var vstatk  []string
	var vstatv []int
	for i:=1;i<27;i++ {
		vstatk = append(vstatk, strconv.Itoa(i))
		vstatv = append(vstatv, i)
	}
	for i := 0; i < len(vstatk); i++ {
		hash[vstatk[i]] = vstatv[i]
	}


	fmt.Println("----=+++++=-----")
	fmt.Println("----=+++++=-----")
	fmt.Println("----=+++++=-----")
}
