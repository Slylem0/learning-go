package main

import (
	"fmt"
)

func Colors() []string {
	color := []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

	return color
}

func ColorCode(color string) int {
	n := Colors()
	for i, c := range n {
		if c == color {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(ColorCode("black"))
}
