package storage

import . "../goods"

type GoodsMap struct {
	Count int
	Goods Goods
}

type Storage struct {
	ID     uint64
	Region int
	Goods  []GoodsMap
}
