package tree

import "fmt"

const cx = 0.0001

func Sqrt(n float64) float64 {
	return SqrtF(n, 1.0)
}

func SqrtF(n, guess float64) float64 {
	if n-guess*guess == 0 {
		return guess
	}
	if n-guess*guess > cx {
		fmt.Println("----", n, guess)
		guess = (n + guess) / 2
		return SqrtF(n, guess)
	}

	return guess
}
