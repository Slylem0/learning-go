package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// devuelve un saludo a una persona especifica
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	mesagge := fmt.Sprintf(randomFormat(), name)
	return mesagge, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message

	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"hola, %v bienvenido",
		"que tal, soy %v",
		"peso deidad saluda a %v",
		"aws sbg manda ping a %v",
	}

	return formats[rand.Intn(len(formats))]
}
