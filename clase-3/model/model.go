package model

import "errors"

var (
	// ErrPersonCanNotBeNil La persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("La persona no puede ser nula")

	// ErrIDPersonDoesNotExists La persona no existe
	ErrIDPersonDoesNotExists = errors.New("La persona no existe")
)
