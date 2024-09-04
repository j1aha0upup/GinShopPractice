package model

type GoodsAttribute struct {
	Id             int
	GoodsId        int
	TypeId         int
	AttributeId    int
	AttributeType  int
	AttributeTitle string
	AttributeValue string
	AddTime        int
	Status         int
	Sort           int
}

func (goodsAttribute *GoodsAttribute) TableName() string {
	return "goods_attribute"
}
