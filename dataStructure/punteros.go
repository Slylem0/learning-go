package main

import (
	"fmt"
)

type personas struct {
	nombre string
	edad   int
	correo string
}

func (p *personas) saludarPersona() {
	fmt.Println("hola mi nombre es: ", p.nombre)
}

func main() {
	var x int = 10
	fmt.Println(x)

	editar(&x)
	fmt.Println(x)

	//punteros de persona
	//
	p := personas{"nicolas marin", 19, "pnicolasmarein@gmail.com"}

	p.saludarPersona()
}

func editar(x *int) {
	*x = 20
}
