package main

import "POO/book"

func main() {
	myBook, _ := book.NewBook("la soledad de los numeros primos", "paulinni grosephi", 167)

	myBook.PrintInfo()
}
