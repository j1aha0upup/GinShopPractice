package model

type Focus struct {
	Id            int
	Title         string
	Type          int //1为网页端轮播图2为小程序端轮播图3为APP端轮播图
	FocusImg      string
	Link          string
	Sort          int
	Status        int
	AddTime       int
	FocusImgSmall string
}

func (focus *Focus) TableName() string {
	return "focus"
}
