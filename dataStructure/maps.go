package main

import (
	"fmt"
)

func main() {
	colors := map[string]int{
		"amarillo": 0o0000,
		"rojo":     1111111,
		"verde":    222222222,
	}

	fmt.Println(colors)
	fmt.Println(colors["rojo"])

	colors["negro"] = 12345444444444

	fmt.Println(colors)

	valor, ok := colors["rojo"]
	fmt.Println(valor, ok)

	valor1, ok1 := colors["blanco"]
	fmt.Println(valor1, ok1)

	valor2 := colors["jajhaja"]
	fmt.Println(valor2)

	delete(colors, "azul")
	fmt.Println(colors)

	for clave, valor := range colors {
		fmt.Println(clave, valor)
	}
}
