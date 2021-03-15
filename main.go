package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registrando la ruta y el handler(saludar)
	http.HandleFunc("/saludar", saludar)
	http.HandleFunc("/despedir", despedir)
	http.ListenAndServe(":8080", nil)
}

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func despedir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hasta luego")
}
