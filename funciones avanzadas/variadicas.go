package main

import (
	"fmt"
)

func suma(name string, nums ...int) {
	fmt.Printf("hola %s, la suma es: %T - %v\n", name, nums, nums)
	var total int
	for _, i := range nums {
		total += i
	}
	fmt.Println(total)
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	suma("nicolas", 34, 7, 23)
	suma("paula", 2, 3, 4, 5, 6, 7, 8, 10001)

	fmt.Println()

	imprimirDatos("hola", 1, 3, 9.777, "santiago", true)
}
