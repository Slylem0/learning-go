package main

import (
	"fmt"

	"rsc.io/quote"
)

// funcion que da un hola mundo
func saludar() {
	fmt.Println("hola mundo desde go ")
	fmt.Println(quote.Hello())
}

//para ejecutar archivos de run hay que hacer un go run que es el que compila el archivo
//
//o tambien el go build para hacer un ejecutable como ta
