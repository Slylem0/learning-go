package main

import (
	"fmt"
)

func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else if kind == "track" {
		return true
	} else {
		return false
	}
}

func ChooseVehicle(option1, option2 string) bool {
	if option1 > option2 {
		return true
	} else {
		return false
	}
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
		return originalPrice - originalPrice*0.2
	} else if age >= 10 {
		return originalPrice / 2
	} else {
		return originalPrice - originalPrice*0.3
	}
}

func main() {
	a := CalculateResellPrice(1000, 2)
	fmt.Println(a)
}
