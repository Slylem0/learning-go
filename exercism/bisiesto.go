package main

import (
	"fmt"
)

func main() {
	fmt.Println("hola mundo")
	fmt.Println(bisiesto(2100))
}

func bisiesto(year int) bool {
	var a, b bool

	if year%4 == 0 {
		a = true
	}

	if year%100 == 0 {
		a = true
	}

	if a && !b {
		return true
	}

	if a && b && year%4 == 0 {
		return true
	}

	return false
}
