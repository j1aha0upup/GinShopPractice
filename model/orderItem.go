package model

type OrderItem struct {
	Id           int
	Uid          int
	OrderId      int
	GoodsId      int
	GoodsTitle   string
	GoodsImage   string
	GoodsPrice   float64
	AddTime      int
	GoodsNum     int
	GoodsVersion string
	GoodsColor   string
}

func (orderItem *OrderItem) TableName() string {
	return "order_item"
}
