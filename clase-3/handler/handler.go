package handler

import "github.com/sjaureguio/golang-api/clase-3/model"

type Storage interface {
	Create(person *model.Person) error
	Update(ID uint, person *model.Person) error
	Delete(ID uint) error
	GetByID(ID uint) error
	GetAll() (model.Persons, error)
}
