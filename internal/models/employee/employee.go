package employee

import . "../storage"
import . "../users"

type Position struct {
	name string
}

type Employee struct {
	ID         uint64
	position   Position
	firstName  string
	lastName   string
	patronymic string
	user       User
	storage    Storage
}
