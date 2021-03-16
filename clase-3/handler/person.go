package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/sjaureguio/golang-api/clase-3/model"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una structura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {

		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Success, "Persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Success, "Ok", data)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Atoi convierte un string a entero
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El ID debe ser un n° entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtner las personas", nil)
		responseJSON(w, http.StatusInternalServerError, response)

		return
	}

	response := newResponse(Success, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Atoi convierte un string a entero
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newResponse(Error, "El ID debe ser un n° entero positivo", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de persona no existe", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := newResponse(Error, "Ocurrió un error al intentar eliminar", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Error, "Ok", nil)
	responseJSON(w, http.StatusOK, response)
}
