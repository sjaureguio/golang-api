package main

import (
	"net/http"
	// "github.com/sjaureguio/golang-api/clase-3/handler"
	"github.com/sjaureguio/golang-api/clase-3/storage"
)

func main() {
	store := storage.NewMemery()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)

	http.ListenAndServe(":8080", mux)
}
