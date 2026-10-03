package main

import (
	"fmt"
)

func main() {
	saludo := func(name string) {
		fmt.Printf("hola, %s\n", name)
	}

	saludo("nicolas")
	nombres := []string{"pablo", "laura", "abril", "juan", "manuel"}

	for _, name := range nombres {
		saludo(name)
	}
}
