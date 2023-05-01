package tasks

import (
	. "../employee"
	. "../orders"
	"time"
)

type TaskType string

const (
	acceptGoods TaskType = "ACCEPT_GOODS"
	collect     TaskType = "COLLECT"
	pack        TaskType = "PACK"
	otherAction TaskType = "OTHER_ACTION"
)

type TaskStatus string

const (
	created    TaskStatus = "CREATED"
	inProgress TaskStatus = "IN_PROGRESS"
	completed  TaskStatus = "COMPLETED"
)

type Task struct {
	ID                 uint64
	TypeName           TaskType
	Status             TaskStatus
	Description        string
	OrderingManager    Employee
	ProcessingEmployee Employee
	OrderingDate       time.Time
	DeadlineDate       time.Time
	EndDate            time.Time
	Order              Order
}
