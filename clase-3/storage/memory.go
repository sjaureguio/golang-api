package storage

import (
	"fmt"

	"github.com/sjaureguio/golang-api/clase-3/model"
)

// Memory .
type Memory struct {
	currentID uint
	Persons   map[uint]model.Person
}

// NewMemery return a instance of Memory
func NewMemery() Memory {
	persons := make(map[uint]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create is used for create new person
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update .
func (m *Memory) Update(ID uint, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete borra de la memoria la persona
func (m *Memory) Delete(ID uint) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memory) GetByID(ID uint) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return person, nil
}

// GetAll retorna todas las personas que están en la memoria
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
