package handler

import (
	"net/http"
)

func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
}
