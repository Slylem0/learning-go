package greetings

import "fmt"

// devuelve un saludo a una persona especifica
func hello(name string) string {
	mesagge := fmt.Sprintf("hola %v un gran saludo", name)
	return mesagge
}
