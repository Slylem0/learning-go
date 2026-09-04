package main

import (
	"fmt"
)

// resulta que en go no hay ciclos while se hace todo desde
// el ciclo for
func main() {
	// bucle for incializado en si mismo
	for i := 1; i <= 10; {
		fmt.Println(i)
		i++
	}

	fmt.Println("______________________________________")

	for j := 1; j <= 10; j++ {

		if j == 5 {
			continue
		}
		fmt.Println(j)
	}
}
