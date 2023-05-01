package orders

import . "../customer"
import . "../goods"

type OrderStatus struct {
	ID   uint64
	Name string
}

type OrderGoods struct {
	Goods Goods
	Count int
}

type Order struct {
	ID          uint64
	OrderStatus OrderStatus
	Customer    Customer
	Card        []OrderGoods
}
