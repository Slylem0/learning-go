package main

import (
	"errors"
	"fmt"
)

func dividir(numerador, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se peude dividir entre 0 amigue")
	}

	return numerador / divisor, nil
}

func main() {
	resultado, err := dividir(10, 0)
	fmt.Println(resultado, err)

	fmt.Println(dividir(10, 9))
}
