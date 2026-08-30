package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 50
	var integer32 int32 = 100

	i := "100"
	h, _ := strconv.Atoi(i)

	fmt.Println(h + h)

	fmt.Println(int32(integer16) + integer32)

	//funcio IOTA integnet to Ascci
	//
	n := 42
	j := strconv.Itoa(n)
	fmt.Println(j, " ", j)
}
