package main

import (
	"fmt"
	"log"

	"github.com/Slylem0/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"alex", "memento", "gfustavo", "pablo"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// message, err := greetings.Hello("moon")
	// if err != nil {
	//	log.Fatal(err)
	// }

	fmt.Println(messages)
}
