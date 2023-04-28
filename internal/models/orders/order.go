package orders

import . "../customer"
import . "../goods"

type OrderStatus struct {
	ID   uint64
	name string
}

type OrderGoods struct {
	goods Goods
	count int
}

type Order struct {
	ID          uint64
	orderStatus OrderStatus
	customer    Customer
	card        []OrderGoods
}
