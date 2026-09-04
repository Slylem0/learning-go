package main

import (
	"fmt"
)

func main() {
	hello("pablo nicolas la monda")
	fmt.Println(hi("pablitogamer32"))
	calc(9, 9)
	sum, mul := calc(8, 9)
	fmt.Println(sum, mul)
}

func hello(name string) {
	fmt.Println("hola", name)
}

func hi(name string) string {
	return "hola, " + name
}

func calc(a, b int) (int, int) {
	suma := a + b
	mult := a * b
	return suma, mult
}
