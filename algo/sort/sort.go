package sort

import "fmt"

func MidSort() {

}

/**
Input: {5 2 4 6 1 3}。
首先拿起第一张牌, 手上有 {5}。
拿起第二张牌 2, 把 2 insert 到手上的牌 {5}, 得到 {2 5}。
拿起第三张牌 4, 把 4 insert 到手上的牌 {2 5}, 得到 {2 4 5}。
*/
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i] //当前牌
		j := i - 1
		for arr[j] > tmp && j >= 0 {
			arr[j+1] = arr[j]
			j--
		}
		if j+1 != i {
			arr[j+1] = tmp
		}
	}
}

func InsertSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		tmp := values[i] // 从第二个数开始，从左向右依次取数
		key := i - 1     // 下标从0开始，依次从左向右

		// 每次取到的数都跟左侧的数做比较，如果左侧的数比取的数大，就将左侧的数右移一位，直至左侧没有数字比取的数大为止
		for key >= 0 && tmp < values[key] {
			values[key+1] = values[key]
			key--
			//fmt.Println(values)
		}

		// 将取到的数插入到不小于左侧数的位置
		if key+1 != i {
			values[key+1] = tmp
		}
		//fmt.Println(values)
	}
}

/**
状态存储
*/
func Fb(n int) {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	fmt.Println(arr)

}

type Sort interface {
	Sort(arr []int) []int
}

/**
选择排序，每次选出数组中的最大（小值）
*/
func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		max := i // 假设当前最大
		for j := i; j < length; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		if max != i {
			arr[i], arr[max] = arr[max], arr[i] // 交换
		}
	}
	return arr
}

func HeapSortMax(arr []int ,length int) []int {
	for i := 0; i < length; i++ {
		
	}
	return arr
}
