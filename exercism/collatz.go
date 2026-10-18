package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	acc := 0
	er := errors.New("no se peude con negativos o 0")

	if n >= 0 {
		return 0, er
	}

	for {
		if n == 1 {
			return acc, nil
		} else if n%2 == 0 {
			n = n / 2
			acc += 1
		} else {
			n = n*3 + 1
			acc += 1
		}
	}
}

func main() {
	fmt.Println(CollatzConjecture(10))
}
