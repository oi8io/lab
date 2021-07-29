package diu

import (
	"fmt"
	"testing"
)

/*
map[0:      1:false 2:false 3:false 4:false 5:false 6:false 7:false 8:false]
map[0:false 1:false 2:false 3:false 4:false 5:false 6:false 7:      8:false]
map[0:false 1:      2:false 3:false 4:false 5:false 6:false 7:false 8:false]
map[0:false 1:false 2:false 3:false 4:false 5:false 6:false 7:false 8:     ]
map[0:false 1:false 2:false 3:false 4:false 5:      6:false 7:false 8:false]
map[0:false 1:false 2:      3:false 4:false 5:false 6:false 7:false 8:false]
map[0:false 1:false 2:false 3:false 4:      5:false 6:false 7:false 8:false]
map[0:false 1:false 2:false 3:false 4:false 5:false 6:false 7:false 8:false]
map[0:false 1:false 2:false 3:      4:false 5:false 6:false 7:false 8:false]

*/

func Smart() {
	rows_X := make(map[int]bool)
	rows_Y := make(map[int]bool)
	rows_P := make(map[int]bool)

	for x := 0; x < 8; x++ {
		if rows_X[x] {
			continue
		}
		for y := 0; y < 8; y++ {
			if rows_Y[y] || rows_Y[y+1]{ //坐标被占用
				continue
			}
			if y >= 1 && rows_Y[y-1] {
				continue
			}
			if rows_P[(x-1)*10+(y-1)] {
				continue
			}
			rows_X[x] =true
			rows_Y[y] = true
			rows_P[x*10+y] = true
			break
		}
	}
	fmt.Println(rows_P)
}
func queen(a [8]int, cur int) {
	if cur == len(a) {
		//fmt.Print(a)
		//fmt.Println()
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1)
		}
	}
}
func TestSmart(t *testing.T) {
	var balance = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	queen(balance, 0)
}
