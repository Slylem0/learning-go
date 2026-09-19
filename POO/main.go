package main

import (
	"POO/animal"
	"POO/book"
	"fmt"
)

func main() {
	myBook, _ := book.NewBook("la soledad de los numeros primos", "paulinni grosephi", 167)

	myBook.PrintInfo()
	myBook.Settitle("cien años de soledad modified")

	fmt.Println(myBook.Gettitle())

	myBook.PrintInfo()

	myTextbook := book.NewtextBook("la cartilla nacho", "cifuentes",
		100, "memento farsa", "2")

	myTextbook.PrintInfo()

	//myBook.PrintInfo()
	//myTextbook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextbook)

	///////////////////////////////////////////////////////////////////////////////////
	//
	//
	//
	miPerro := animal.Perro{Nombre: "pepe"}
	migato := animal.Gato{Nombre: "nacho"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&migato)

	animales := []animal.Animal{
		&animal.Gato{Nombre: "pepe"},
		&animal.Perro{Nombre: "etesech"},
		&animal.Gato{Nombre: "daniel"},
		&animal.Perro{Nombre: "kevin"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
