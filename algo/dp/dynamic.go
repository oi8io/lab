package dp

import "fmt"

/**
27块钱
手里有若干面值5元 7元 2元若干，求最少硬币
f(27) = min{f(27-2)+1,f(27-5)+1,f(27-7)+1}
*/
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func Recursion(x int) int {
	var total = MaxInt
	if x == 0 {
		return 0
	}

	if x == 1 {
		fmt.Println("errr")
		return MaxInt
	}

	//fmt.Println(x)
	if x >= 2 {
		total = Min(Recursion(x-2)+1, total)
	}
	if x >= 5 {
		total = Min(Recursion(x-5)+1, total)
	}
	if x >= 7 {
		total = Min(Recursion(x-7)+1, total)
	}
	return total
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CoinChange(coins []int, amount int) int {
	var minCoinArr = make([]int, amount+1)
	if amount == 0 {
		return 0
	}
	for i := 0; i < amount+1; i++ {
		x := 0
		minCoinArr[i] = MaxInt
		minCoinArr[0] = 0
		for _, c := range coins {
			if i-c >= 0 {
				ret := 1 + minCoinArr[i-c]
				if ret < 0 {
					continue
				}
				if ret > minCoinArr[i] {
					continue
				}
				x = c
				minCoinArr[i] = ret
			}
		}
		fmt.Println(i, "投入：", x)
	}

	if minCoinArr[amount] == MaxInt {
		return -1
	}
	return minCoinArr[amount]
}

/**
有一个国家发现了5座金矿，每座金矿的黄金储量不同，需要参与挖掘的工人数也不同。参与挖矿工人的总数是10人。每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。要求用程序求解出，要想得到尽可能多的黄金，应该选择挖取哪几座金矿？
1、400金/5人   2、500金/5人   3、200金/3人    4、300金/4人    5、350金/3人

*/
func Gold(source map[int]int, people int) int {
	var maxCoinArr = make([]int, people+1)
	var storage = make(map[int]map[int]int)
	for i := 0; i <= people; i++ {
		maxCoinArr[i] = 0
		storage[i] = map[int]int{}
		var x int
		for g, p := range source {
			if i >= p {
				c := maxCoinArr[i-p] + g
				if c > maxCoinArr[i] {
					maxCoinArr[i] = c
					x = g
				}
			}
		}
		if x > 0 {
			fmt.Println(i, x)
			fmt.Println(source[x])
			storage[i] = storage[i-source[x]]
			storage[i][x] = storage[i][x] + 1
			fmt.Println("zhiq", storage[i-source[x]])
			fmt.Println(i, "个工人", "wa", x, maxCoinArr[i])

		}
	}
	for i, i2 := range storage {
		fmt.Println(i, "=>", i2)
	}
	return maxCoinArr[people]
}
