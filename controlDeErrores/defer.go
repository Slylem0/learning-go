package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.Write([]byte("hola entonces si se creo este archivo."))
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}
