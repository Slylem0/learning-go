package main

import (
	"POO/book"
	"fmt"
)

func main() {
	myBook, _ := book.NewBook("la soledad de los numeros primos", "paulinni grosephi", 167)

	myBook.PrintInfo()
	myBook.Settitle("cien años de soledad modified")

	fmt.Println(myBook.Gettitle())

	myBook.PrintInfo()
}
