package main

import (
	"fmt"
)

func dividir(numerador, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(divisor)

	fmt.Println("El resultado es: ", numerador/divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("no se puede dividir por cero")
	}
}

func main() {
	dividir(10, 9)
	dividir(10, 5)
	dividir(10, 0)
}
