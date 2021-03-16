package handler

import (
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/records", h.getAll)
	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
}
