package main

import (
	"log"
	"net/http"

	"github.com/sjaureguio/golang-api/clase-3/authorization"
	"github.com/sjaureguio/golang-api/clase-3/handler"
	"github.com/sjaureguio/golang-api/clase-3/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("No se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemery()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v \n", err)
	}

	// Formato para fecha y hora
	// time.Now().Format("2006-01-02 15:04:05")

	// Para generar un certificado privado
	// openssl genrsa -out app.rsa 1024

	// Para generar un certificado publico
	// openssl rsa -in app.rsa -pubout > app.rsa.pub

	// Ruta cmd
	// C:\WINDOWS\System32\cmd.exe
}
