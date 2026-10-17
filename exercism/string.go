package main

import (
	"fmt"
	"strings"
)

func saludar(custumer string) string {
	const a = "Welcome to the Tech Palace, "
	b := strings.ToUpper(custumer)

	return (a + b)
}

func AddBorder(welcomeMsg string, number int) string {
	a := strings.Repeat("*", number)

	return (a + "\n" + welcomeMsg + "\n" + a)
}

func CleanupMessage(oldMsg string) string {
	limpio := strings.ReplaceAll(oldMsg, "*", "")

	limpio = strings.TrimSpace(limpio)
	return limpio
}

func main() {
	fmt.Println(saludar("kevin"))
	fmt.Println(AddBorder("welcome!", 10))
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`

	fmt.Println(CleanupMessage(message))
}
