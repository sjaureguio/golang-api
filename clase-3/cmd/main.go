package main

import (
	"log"
	"net/http"

	"github.com/sjaureguio/golang-api/clase-3/handler"
	"github.com/sjaureguio/golang-api/clase-3/storage"
)

func main() {
	store := storage.NewMemery()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v \n", err)
	}

	// Formato para fecha y hora
	// time.Now().Format("2006-01-02 15:04:05")
}
