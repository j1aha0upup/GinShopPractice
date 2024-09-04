package model

type GoodsType struct {
	Id          int
	Title       string
	Description string
	AddTime     int
	Status      int
}

func (goodsType *GoodsType) TableName() string {
	return "goods_type"
}
