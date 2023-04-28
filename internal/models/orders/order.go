package orders

import . "../customer"

type OrderStatus struct {
	ID   uint64
	name string
}

type Order struct {
	ID          uint64
	orderStatus OrderStatus
	customer    Customer
	goods       []Goods
}
