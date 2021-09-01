package main

func plus() func() int {
	var x int
	return func() int {
		return x + 1
	}
}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func Fib() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}


