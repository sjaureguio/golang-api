package model

// Community estructura de una comunidad
type Community struct {
	Name string `json:"name"`
}

// Communities slice de
type Communities []Community

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

// Persons slice de personas
type Persons []Person
