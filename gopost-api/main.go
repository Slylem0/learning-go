package main

import (
	"fmt"
	"net/http"
)

// handler
func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "metodo http: %s \n", r.Method)
	fmt.Fprintf(w, "Ruta solicitada: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Host: %s \n", r.Host)
	fmt.Fprintf(w, "user-agent: %s\n", r.UserAgent())
}

func main() {
	http.HandleFunc("/hola", hola)

	fmt.Println("servidor iniciando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
