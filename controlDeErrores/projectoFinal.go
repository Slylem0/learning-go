package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// creamos la estrcutura de contacto
type contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// func para crear contactos en un archivos json.
func saveContactsToFile(contacts []contact) error {
	file, err := os.Create("contact.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// cargar contactos desde un json
func loadContactsfromFile(contacts *[]contact) error {
	file, err := os.Open("contact.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// slice contacts
	var contacts []contact

	// cargar datos
	err := loadContactsfromFile(&contacts)
	if err != nil {
		fmt.Println("error al cargar los contactos ")
	}

	// crear una isntancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		// mostrar opciones de usuario
		fmt.Println(" === gestor de contactos ===\n",
			"1. Agregar un contacto \n",
			"2. Mostrar todos los contactos\n",
			"3. Salir \n",
			"elige una opcion: ")

		// leer la opcion del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("error al leer la opcion: ", err)
			return
		}

		switch option {
		case 1:
			var c contact
			fmt.Println("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Println("email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Println("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			// agregar un contacto a slice
			contacts = append(contacts, c)

			// agregar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("error al guardar el conectato: ", err)
			}
		case 2:
			fmt.Println("========================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. nombre: %s email: %s telefono: %s \n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("===================================================")

		case 3:
			return

		default:
			fmt.Println("opcion no valida")

		}

	}
}
