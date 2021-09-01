package dp

import (
	"fmt"
	"testing"
)

func TestRecursion(t *testing.T) {
	x := Recursion(27)
	fmt.Println(x)
}

func TestMin(t *testing.T) {
	fmt.Println(Min(18, 198))
}

func TestCoinChange(t *testing.T) {
	var coins = []int{1, 2, 5}
	var amount = 11
	var ret int
	//ret = CoinChange(coins, amount)
	//fmt.Println(ret)
	coins = []int{2, 5, 7}
	amount = 27
	//fmt.Println(amount)
	ret = CoinChange(coins, amount)
	fmt.Println(ret)
	return
	for i := 0; i < 8; i++ {
		ret = CoinChange(coins, i)
		fmt.Println(ret)
	}

}

func TestGold(t *testing.T) {
	//400金/5人   2、500金/5人   3、200金/3人    4、300金/4人    5、350金/3人
	var source = make(map[int]int)
	source[400] = 5
	source[500] = 5
	source[200] = 3
	source[300] = 4
	source[350] = 3
	x := Gold(source, 10)
	fmt.Println(x)

}
