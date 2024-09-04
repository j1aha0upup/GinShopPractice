package model

type GoodsImage struct {
	Id       int
	GoodsId  int
	ImageUrl string
	Sort     int
	ColorId  int
	AddTime  int
	Status   int
}

func (goodsImage *GoodsImage) TableName() string {
	return "goods_image"
}
