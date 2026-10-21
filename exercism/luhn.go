package main

import (
	"fmt"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	a := len(id)
	acc := 0
	if a <= 1 {
		return false
	}
	for i := a - 1; i >= 0; i-- {
		if i%2 == 0 {
			number := int(id[i] - '0')
			if number*2 > 9 {
				acc += number*2 - 9
			}
		} else {
			acc += int(id[i] - '0')
		}
	}
	fmt.Println(acc)

	return acc%10 == 0
}

func main() {
	fmt.Println(Valid("4539 3195 0343 6467"))
}
