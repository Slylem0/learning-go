package main

import (
	"fmt"
)

func main() {
	// los slides son dinamicas
	var a []int

	days := []string{
		"domingo", "lunes", "martes", "miercoles",
		"jueves", "viernes", "sabado",
	}

	univalle := []string{
		"univalle", "universidad del perreo",
		"yandel", "wisin",
	}

	diaRebanada := days[0:5]

	fmt.Println(a)
	fmt.Println(days)
	fmt.Println(diaRebanada)

	fmt.Println(len(days))
	fmt.Println(cap(days))

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	//agregar elementos
	//
	diaRebanada = append(diaRebanada, "no", "hay", "clase")

	diaRebanada = append(diaRebanada, univalle...)
	fmt.Println(diaRebanada)

	// make
	nombre := make([]string, 5, 10)
	nombre[0] = "alex"

	fmt.Println(nombre)

	i := []int{1, 2, 3, 4, 5}
	po := make([]int, 5, 10)

	copy(po, i)
	fmt.Println(po)
}
