package customer

import "time"

type CustomerType string

const (
	individual   CustomerType = "INDIVIDUAL"
	organization CustomerType = "ORGANIZATION"
)

type Customer struct {
	ID           uint64
	customerType CustomerType
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
