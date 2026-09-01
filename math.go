package main

import (
	"fmt"
	"math"
)

func main() {
	// pa que salgan decimal hay que poner .0 a los numeros a y b
	a := 3
	b := 10
	// incrementales
	b++
	b++

	// decrementos
	a--
	b--

	a = a + 5
	b += 5

	fmt.Println(b / a)

	//modulo
	//
	fmt.Println(b % a)
	fmt.Println(math.Pi)
}
