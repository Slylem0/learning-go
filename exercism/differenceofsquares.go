package main

import (
	"fmt"
)

func SquareofSum(n int) int {
	acc := 0
	for i := 1; i <= n; i++ {
		acc += i
	}
	return acc * acc
}

func SumofSquares(n int) int {
	acc := 0
	for i := 1; i <= n; i++ {
		acc += i * i
	}
	return acc
}

func Difference(n int) int {
	a := SquareofSum(n)
	b := SumofSquares(n)

	return a - b
}

func main() {
	fmt.Println(SquareofSum(10))
	fmt.Println(SumofSquares(10))
	fmt.Println(Difference(10))
}
