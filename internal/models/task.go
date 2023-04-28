package models

import (
	. "./employee"
	. "./orders"
	"time"
)

type TaskType struct {
	name string
}

type TaskStatus struct {
	name string
}

type Task struct {
	ID                 uint64
	typeName           TaskType
	status             TaskStatus
	description        string
	orderingManager    Employee
	processingEmployee Employee
	orderingDate       time.Time
	deadlineDate       time.Time
	endDate            time.Time
	order              Order
}
