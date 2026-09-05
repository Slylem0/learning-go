package main

import (
	"fmt"
)

// las matrices se definen en una long fija
func main() {
	var a [5]int

	a[0] = 10
	a[1] = 88

	fmt.Println(a)

	b := [5]int{1 * 33, 2 * 5, 3, 4, 5}
	fmt.Println(b)

	// hacer una matriz sin saber su tamaño final
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	// elemento
	fmt.Println(a[0])

	// recorrer elemtnos con for
	for i := 0; i < len(c); i++ {
		fmt.Print(" ", c[i])
	}

	fmt.Println()
	//version modernizada
	//
	for i := range len(c) {
		fmt.Print(" ", c[i])
	}

	for index, value := range c {
		fmt.Printf("indice: %d, valor %d \n", index, value)
	}

	//definir uan matriz bidimensional
	//
	matrix := [5][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(matrix)
}
