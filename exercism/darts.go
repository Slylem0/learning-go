package main

import (
	"fmt"
	"math"
)

func Score(x, y float64) int {
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if dist <= 1 {
		return 10
	} else if dist <= 5 {
		return 5
	} else if dist <= 10 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(Score(1, 1))
}
