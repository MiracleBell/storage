package delivery

import "time"
import . "../orders"

type DeliveryService struct {
	name string
}

type Delivery struct {
	ID                   uint64
	order                Order
	address              string
	expectedDeliveryDate time.Time
	deliveryService      DeliveryService
}
