package handler

import "net/http"
import "github.com/sjaureguio/golang-api/clase-3/storage"

func RoutePerson(mux *http.ServeMux, storage.Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
}