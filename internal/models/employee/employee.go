package employee

import . ".."
import . "../users"

type Employee struct {
	ID         uint64
	position   Position
	firstName  string
	lastName   string
	patronymic string
	user       User
	storage    Storage
}
