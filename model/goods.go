package model

import (
	"practice_gin_store/util"
	"strconv"
	"strings"
)

type Goods struct {
	Id               int
	Title            string
	SubTitle         string
	GoodsSn          string
	CateId           int
	ClickCount       int
	GoodsNumber      int
	ShopPrice        float64
	MarketPrice      float64
	GoodsRelation    string
	GoodsAttrs       string
	GoodsColor       string
	GoodsVersion     string
	GoodsImage       string
	GoodsImages      string
	GoodsFitting     string
	GoodsKeywords    string
	GoodsDescription string
	GoodsContent     string
	GoodsGift        string
	MoreChoice       string
	GoodsTypeId      int
	IsDelete         int
	IsHot            int
	IsBest           int
	IsNew            int
	Status           int
	Sort             int
	AddTime          int
}

func (goods *Goods) TableName() string {
	return "goods"
}
func IsColorInGoods(goodsColorIdList string, goodsColorId int) bool {
	goodsColorIdStrList := strings.Split(goodsColorIdList, " ")
	for _, v := range goodsColorIdStrList {
		if v == strconv.Itoa(goodsColorId) {
			return true
		}
	}
	return false
}

func AttributeOptionList(AttributeId int) []string {
	goodsTypeAttribute := GoodsTypeAttribute{}
	util.Db.Where("id=?", AttributeId).Find(&goodsTypeAttribute)
	return strings.Split(goodsTypeAttribute.AttributeValue, "\n")
}
