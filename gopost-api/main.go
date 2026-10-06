package main

import (
	"encoding/json"
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

func sipo(w http.ResponseWriter, r *http.Request) {
	datos := map[string]string{
		"mensaje": "hola mundo desde go WAZAAAAAAAAA",
		"status":  "OK",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(datos)
}

func main() {
	http.HandleFunc("/hola", hola)
	http.HandleFunc("/sipo", sipo)
	fmt.Println("servidor iniciando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
