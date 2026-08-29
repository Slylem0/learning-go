package main

import (
	"fmt"
	"math"
	"strconv"
)

func entiendo() {
	fmt.Println("_________________ TIPOS DE DATOS EN GOLANG ___________________")
	// almacena desde -127 hasta 128 hay diferentes tipos de int
	var integer int8 = 127
	var entero uint64 = 7834682796937429687

	var float float64 = 1.87987988712398971823

	// variables string
	fullName := "pablio gamer32 \t( alias \"pablo nicolas\") \n"
	fmt.Println(fullName)

	//hay un paquete llamado math que sirve para operaciones matematicas
	//
	fmt.Println(math.MaxInt64, math.MinInt64)
	fmt.Println(integer, entero, float)

	// tipo de dato byte como asccii
	var a byte = 'a'
	fmt.Println(a)

	// podemos accede a caracrteres especificos en un string pero devuelve en su formato asccii

	s := "holaa"
	fmt.Println(s[0])

	// valorres predeterminados

	var (
		defaultint    int
		defaulutuint  uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultint, defaulutuint, defaultFloat, defaultBool, defaultString)

	//conversion de tipos
	//
	//

	h := "100"
	i, _ := strconv.Atoi(h)

	fmt.Println(i + i)
}

func main() {
	entiendo()
}
