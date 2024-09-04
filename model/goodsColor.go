package model

type GoodsColor struct {
	Id         int
	ColorName  string
	ColorValue string
	Status     int
}

func (goodsColor *GoodsColor) TableName() string {
	return "goods_color"
}
