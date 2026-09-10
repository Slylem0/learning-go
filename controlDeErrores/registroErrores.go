package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("main()")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("oye soy un log")
}
