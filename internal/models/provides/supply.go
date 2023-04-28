package provides

import "time"
import . "../storage"
import . "../goods"

type SupplyMap struct {
	goods Goods
	count int
}

type Supply struct {
	ID                   uint64
	provider             Provider
	expectedDeliveryDate time.Time
	realDeliveryDate     time.Time
	storage              Storage
	detail               []SupplyMap
}
