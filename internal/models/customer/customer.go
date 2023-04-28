package customer

import "time"

type Customer struct {
	ID        uint64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
