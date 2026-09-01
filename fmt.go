package main

import (
	"fmt"
)

func main() {
	var name string
	var edad int

	fmt.Print("digite su nombre por favor: ")
	fmt.Scanln(&name)
	fmt.Print("digite su edad pro favor: ")
	fmt.Scanln(&edad)

	fmt.Printf("hola me llamo %s y tengo %d\n", name, edad)
	fmt.Printf("el tipo de dato de nombre: %T \n  ", name)
	greeting := fmt.Sprintf("hola soy %s y tengo %d", name, edad)

	fmt.Println(greeting)
}
