package model

type GoodsCate struct {
	Id              int
	Title           string
	CateImage       string
	FilterAttribute string
	Link            string
	Pid             int
	Template        string
	SubTitle        string
	Keywords        string
	Description     string
	Sort            int
	AddTime         int
	Status          int
	GoodsCateList   []GoodsCate `gorm:"foreignKey:Pid"`
}

func (goodsCate *GoodsCate) TableName() string {
	return "goods_cate"
}
