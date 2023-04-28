package provides

import "time"
import . ".."

type Supply struct {
	ID                   uint64
	provider             Provider
	expectedDeliveryDate time.Time
	realDeliveryDate     time.Time
	storage              Storage
	goods                []Goods
}
