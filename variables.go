package main

import (
	"fmt"
)

const euler = 2.718281

// varior iota
const (
	domingo = iota + 1
	lunes
	martes
	miercoles
	jueves
	viernes
	sabado
)

// declaracion de variables en golang
func main() {
	//var firstName, lastName string
	//var age int
	//
	//declarar constantes
	const pi = 3.14

	// inicializar variables
	var (
		firstName, lastName string
		age                 int
	)
	// otra forma de declarar variables

	firstName = "pablo"
	lastName = "nicolas marin"
	age = 19

	fmt.Println(firstName, lastName, age)
	fmt.Println(pi, euler)
	fmt.Println(viernes)
}
