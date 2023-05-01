package provides

import "time"
import . "../storage"
import . "../goods"

type SupplyMap struct {
	Goods Goods
	Count int
}

type Supply struct {
	ID                   uint64
	Provider             Provider
	ExpectedDeliveryDate time.Time
	RealDeliveryDate     time.Time
	Storage              Storage
	Detail               []SupplyMap
}
