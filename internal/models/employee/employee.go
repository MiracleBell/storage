package employee

import . "../storage"
import . "../users"

type Position struct {
	name string
}

type Employee struct {
	ID         uint64
	Position   Position
	FirstName  string
	LastName   string
	Patronymic string
	User       User
	Storage    Storage
}
