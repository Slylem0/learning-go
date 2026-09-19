package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

type Gato struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "hace wow wow ")
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + "hace miau miau ")
}

func HacerSonido(animal Animal) {
	animal.Sonido()
}
