package main

import (
	"bufio"
	"fmt"
	"os"
)

type tarea struct {
	name       string
	descr      string
	completado bool
}

type listaTarea struct {
	tareas []tarea
}

func (l *listaTarea) addTarea(t tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para colcoar como completado una tarea
func (l *listaTarea) completeTarea(i int) {
	l.tareas[i].completado = true
}

// metodo para editar una tarea
func (l *listaTarea) editarTarea(index int, t tarea) {
	l.tareas[index] = t
}

// metodo para eliminatar una tarea
// pide un ountero
// y un index
func (l *listaTarea) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	lista := listaTarea{}

	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("selecione la opcion\n",
			"1. agregar tarea \n",
			"2. marcar tarea como completada \n",
			"3. editar tarea \n",
			"4. eliminar tarea \n",
			"5. salir \n")
		fmt.Println("ingrese una opcion por favor: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t tarea
			fmt.Println("ingrese el nombre de la tarea")
			t.name, _ = leer.ReadString('\n')
			fmt.Println("ingrese la descripcion de la tarea")
			t.descr, _ = leer.ReadString('\n')
			lista.addTarea(t)
			fmt.Println("tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("por favir ingrese el indice")
			fmt.Scanln(&index)
			lista.completeTarea(index)

		case 3:
			var index int
			var t tarea
			fmt.Println("ingrese el indice que quiere editar la tarea")
			fmt.Scanln(&index)
			fmt.Println("ingese el nombre de la tarea")
			t.name, _ = leer.ReadString('\n')
			fmt.Println("ingrese la descripcion de la tarea")
			t.descr, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
		case 4:
			var index int
			fmt.Println("ingrese el indice de la tarea que quiere eliminar")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("tarera elinianda satisfactoriamente")
		case 5:
			fmt.Println("saliendo del programa...")
			return
		default:
			fmt.Println("opcion invalida")

		}
		fmt.Println("LISTA DE TAREAS")
		fmt.Println("____________________________________________________")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - completado: %t \n", i, t.name, t.descr, t.completado)
		}
	}
}
