package model

type Cart struct {
	Id           int
	Title        string
	Price        float64
	GoodsVersion string
	Uid          int
	Num          int
	GoodsGift    string
	GoodsFitting string
	GoodsColor   string
	GoodsImage   string
	GoodsAttr    string
	GoodsId      int
	Checked      bool
}

func (cart *Cart) TableName() string {
	return "cart"
}
