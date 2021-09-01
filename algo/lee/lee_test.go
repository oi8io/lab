package Lee

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

func TestUniquePaths(t *testing.T) {
	var x, y = 3, 4
	UniquePaths(x, y)
}

func TestCanJump(t *testing.T) {
	var nums []int
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	tr := CanJump(nums)
	fmt.Println(tr)
	nums = []int{3, 2, 1, 0, 4}
	f := CanJump(nums)
	fmt.Println(f)

}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	arr := TwoSum(nums, target)
	fmt.Println(arr)

}

func TestAddTwoNumbers(t *testing.T) {
	type cases struct {
		a    []int
		b    []int
		want int
	}
	var inputs = []cases{
		{
			a:    []int{9, 8},
			b:    []int{1},
			want: 90,
		},
		{
			a:    []int{2, 4, 3},
			b:    []int{5, 6, 4},
			want: 807,
		},
		{
			a:    []int{9, 9},
			b:    []int{1},
			want: 100,
		},
	}
	for _, i2 := range inputs {
		if x := AWN(i2.a, i2.b); x != i2.want {
			t.Errorf("[%+v,%+v] want %d but %d ", i2.a, i2.b, i2.want, x)
		}
	}
}

func TestConvertToBase7(t *testing.T) {

	inputs := []struct {
		num  int
		want string
	}{
		{
			num:  7,
			want: "10",
		}, {
			num:  -7,
			want: "-10",
		},
	}

	for _, x := range inputs {
		if s := convertToBase7(x.num); s != x.want {
			t.Errorf(" %d want %s but %s ", x.num, x.want, s)
		}
	}

}

func TestFindMinStep(t *testing.T) {
	inputs := []struct {
		board string
		hand  string
		want  int
	}{
		{
			board: "WRRBBW",
			hand:  "RB",
			want:  -1,
		}, {
			board: "WWRRBBWW",
			hand:  "WRBRW",
			want:  2,
		}, {
			board: "G",
			hand:  "GGGGG",
			want:  2,
		},
	}

	for _, x := range inputs {
		if s := findMinStep(x.board, x.hand); s != x.want {
			t.Errorf(" [%s,%s] want %d but %d ", x.board, x.hand, x.want, s)
		}
	}
}

func TestReverseList(t *testing.T) {
	a := NewListNode([]int{1, 2, 3, 4, 5})
	//fmt.Println(a.Next)
	b := reverseList(a)
	for {
		if b == nil {
			break
		}
		fmt.Println(b.Val)
		b = b.Next
	}

}

func TestMergeTwoLists(t *testing.T) {
	type cases struct {
		a    []int
		b    []int
		want int
	}
	var inputs = []cases{
		{
			a:    []int{1, 2, 4},
			b:    []int{1, 3, 4},
			want: 443211,
		}, {
			a:    []int{1, 5, 9},
			b:    []int{2, 3, 4, 7, 8},
			want: 12345789,
		},
	}
	for _, i2 := range inputs {
		if x := MWN(i2.a, i2.b); x != i2.want {
			t.Errorf("[%+v,%+v] want %d but %d ", i2.a, i2.b, i2.want, x)
		}
	}
}

func TestFiveKM(t *testing.T) {
	x, y := FiveKM(500)
	fmt.Println(x)
	fmt.Println(y)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	inputs := []struct {
		s    string
		want int
	}{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {" ", 1}, {"au", 2}, {"dvdf", 3}}
	for _, i2 := range inputs {
		if x := lengthOfLongestSubstring(i2.s); x != i2.want {
			t.Errorf("%s want %d but %d ", i2.s, i2.want, x)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	inputs := []struct {
		s    string
		want string
	}{{"abcabcbb", "bcb"}, {"babad", "bab"}, {"cbbd", "bb"}}
	for _, i2 := range inputs {
		if x := longestPalindrome(i2.s); x != i2.want {
			t.Errorf("%s want %s but %s ", i2.s, i2.want, x)
		}
	}
}

func TestLongestPalindromeX(t *testing.T) {
	inputs := []struct {
		s    string
		want string
	}{{"abcabcbb", "bb"}, {"babad", "bab"}, {"cbbd", "bb"}, {"cbb", "bb"}}
	for _, i2 := range inputs {
		if x := checkPalindrome(i2.s); x != i2.want {
			t.Errorf("%s want %s but %s ", i2.s, i2.want, x)
		}
	}
}

func TestAWN(t *testing.T) {
	inputs := []struct {
		s    string
		want bool
	}{{"aba", true}, {"aaa", true}, {"abcabcbb", false}, {"abb", false}, {"a", true}}
	for _, i2 := range inputs {
		if x := isPalindrome(i2.s); x != i2.want {
			t.Errorf("%s want %v but %v ", i2.s, i2.want, x)
		}
	}
}

func TestGold2(t *testing.T) {
	inputs := []struct {
		s    string
		want string
	}{{"aba", "aba"},
		//{"abcabcbb", "bcb"}, {"babad", "bab"}, {"cbbd", "bb"}, {"cbb", "bb"}
	}
	for _, i2 := range inputs {
		if x := lastPalindrome(i2.s); x != i2.want {
			t.Errorf("%s want %s but %s ", i2.s, i2.want, x)
		}
	}
}

func TestGetFirstChar(t *testing.T) {
	inputs := []struct {
		s    string
		want uint8
	}{{"leetcode", 'l'}, {"cc", 'c'},
		//{"abcabcbb", "bcb"}, {"babad", "bab"}, {"cbbd", "bb"}, {"cbb", "bb"}
	}
	for _, i2 := range inputs {
		if x := GetFirstChar(i2.s); x != i2.want {
			t.Errorf("%s want %d but %d ", i2.s, i2.want, x)
		}
	}
}

func TestHanoi(t *testing.T) {
	sum:= Hanoi(5, "A", "B","C")
	fmt.Println(sum)

	x := FactTail(5)
	fmt.Println(x)
}
