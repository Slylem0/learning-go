package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Println("WELCOME TO THIS CALCULATOR, WE NEED THIS INFROMATION")
	fmt.Print("ingrese el cateto a: ")
	fmt.Scanln(&a)

	fmt.Print("ingresa el cateto b: ")
	fmt.Scanln(&b)

	// ya con esto vamos a calcular la hipotenusa del triangulo
	h := math.Sqrt((math.Pow(a, 2) + math.Pow(b, 2)))

	fmt.Printf("la hipotenusa de este triangulo es: %f \n", h)

	area := (a * b) / 2
	fmt.Printf("la base del triangulo es: %.2f \n", area)

	perimetro := a + b + h
	fmt.Printf("El perimetro del triangulo es: %.2f", perimetro)
}
