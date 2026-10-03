package main

import (
	"fmt"
)

func incrementar() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	number := incrementar()
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
}
