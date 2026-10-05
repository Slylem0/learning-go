package main

import (
	"fmt"
)

func printList(list ...any) {
	for _, i := range list {
		fmt.Println(i)
	}
}

func main() {
	printList("maria", 1, 2, 8.67465, true, 999)
}
