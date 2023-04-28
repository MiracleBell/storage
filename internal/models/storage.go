package models

type GoodsMap struct {
	count int
	goods Goods
}

type Storage struct {
	ID     uint64
	region int
	goods  GoodsMap
}
