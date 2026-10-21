package main

import (
	"fmt"
	"strings"
)

func IsValidISBN(isbn string) bool {
	// entry a string
	isbn = strings.ReplaceAll(isbn, " ", "")
	isbn = strings.ReplaceAll(isbn, "-", "")
	acc := 0
	numeros := [...]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// cadena ya formateada y limpia
	if len(isbn) != 10 {
		return false
	}
	for i := 0; i < len(isbn); i++ {
		if isbn[i] == 'X' && i == 9 {
			acc += 10 * numeros[i]
		} else if isbn[i] >= '0' && isbn[i] <= '9' {
			acc += int(isbn[i]-'0') * numeros[i]
		} else {
			return false
		}
	}
	return acc%11 == 0
}

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}
