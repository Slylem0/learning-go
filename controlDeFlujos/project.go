package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// crear un numero aleatorio
	play()
}

func play() {
	numberRandom := rand.Intn(100) + 1
	var numeroIngresado, intentos int

	const intentosMaximos = 10

	for intentos < intentosMaximos {
		intentos++
		// le pedios al user que ingrese el numero
		fmt.Printf("ingrese un numero: (intentos dispoibles: %d) ", intentosMaximos-intentos+1)
		fmt.Scanln(&numeroIngresado)

		if numeroIngresado == numberRandom {
			fmt.Println("felicidades ha ganado el juego")
			volverJugar()
		} else if numeroIngresado < numberRandom {
			fmt.Println("el numero que seleccionaste es menor")
		} else if numeroIngresado > numberRandom {
			fmt.Println("el numero seleccionado es mayor")
		}

	}

	fmt.Println("se acabaron los intentos, el numero era: ", numberRandom)
	volverJugar()
}

// ahora creamos una nueva funcion para ver si el usuario
// quiere lver a jugar de nueva
func volverJugar() {
	var opcion string
	fmt.Println("quieres voler a jugar y/n")
	fmt.Scanln(&opcion)

	switch opcion {

	case "y":
		play()

	case "n":
		fmt.Println("gracias por jugar")
	default:
		fmt.Println("seleccione uan opcion valida amigue :) ")
		volverJugar()
	}
}
