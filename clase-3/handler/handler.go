package handler

import "github.com/sjaureguio/golang-api/clase-3/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	// GetByID(ID int) error
	GetAll() (model.Persons, error)
}
