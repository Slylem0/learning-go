package main

import (
	"fmt"
)

type persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p persona
	p.nombre = "pablo nicolas"
	p.edad = 19
	p.correo = "pablo.nicolas@correounivalle.de.co"

	fmt.Println(p)

	persona2 := persona{"camilo lopez", 29, "camilolopez@gmail.com"}

	fmt.Println(persona2)

	persona2.nombre = "nikolas rivera "

	fmt.Println(persona2)
}
