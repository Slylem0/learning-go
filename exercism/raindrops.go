package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("holaa")
	fmt.Println(rain(34))
	fmt.Println("hola")
}

func rain(a int) string {
	var b string
	if a%3 == 0 {
		b = "Pling"
	}

	if a%5 == 0 {
		b += "Plang"
	}

	if a%7 == 0 {
		b += "Plong"
	}

	if b == "" {
		return strconv.Itoa(a)
	}

	return b
}
