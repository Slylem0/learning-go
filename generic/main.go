package main

import (
	"fmt"
)

func printList(list ...any) {
	for _, i := range list {
		fmt.Println(i)
	}
}

func suma[T int | float64](list ...T) T {
	var total T
	for _, i := range list {
		total += i
	}
	return total
}

func includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

type product[T uint | string] struct {
	id    T
	desc  string
	price float32
}

func main() {
	printList("maria", 1, 2, 8.67465, true, 999)

	println("____________________________")
	fmt.Println(suma(1, 2, 3, 4, 5, 6, 7))

	fmt.Println(suma(1.2, 1.3, 1.4, 1.6, 9))

	println("hola mundo en go ")

	slice1 := []string{"a", "b", "c", "d"}
	slice2 := []int{1, 2, 3, 4}

	fmt.Println(includes(slice1, "a"))
	println(includes(slice2, 10))

	product1 := product[uint]{1, "zapatos", 50}
	fmt.Println(product1)

	product2 := product[string]{"wazaaa", "camisas", 100}

	fmt.Println(product2)
}
