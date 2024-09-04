package model

type GoodsTypeAttribute struct {
	Id             int
	Title          string
	AddTime        int
	Status         int
	Sort           int
	AttributeType  int
	AttributeValue string
	TypeId         int
}

func (goodsTypeAttribute *GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
