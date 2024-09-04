package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
	BaseController
}

func (shopController ShopController) Index(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)

	ctx.HTML(http.StatusOK, "shop/index/index.html", gin.H{
		"topNav":    topNav,
		"topNavLen": len(topNav) - 1,
		"middleNav": middleNav,
		"footerNav": footerNav,
		"goodsCate": goodsCate,
		"focus":     focus,
		"isLogin":   user.UserName != "",
		"user":      user,
	})
}

func (shopController ShopController) Category(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	CategoryGoodsList := []model.Goods{}
	categoryId := ctx.Param("categoryId")
	pageNum, _ := strconv.Atoi(ctx.Request.FormValue("page"))
	if pageNum != 0 {
		pageNum--
	}
	pagination := model.Pagination{}
	pagination.PageSize = 5
	pagination.VisiblePages = 5
	pagination.CurrentPage = pageNum + 1

	if tempNavStr, err := util.Redis.GetRedisStringValue("CategoryGoodsList"); err != nil && tempNavStr != "" {
		json.Unmarshal([]byte(tempNavStr), &CategoryGoodsList)
	} else {
		util.Db.Where("cate_id=?", categoryId).Find(&CategoryGoodsList)
		temp, err := json.Marshal(CategoryGoodsList)
		if err != nil {
			fmt.Println("CategoryGoodsList Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("CategoryGoodsList", string(temp), 60)
		}
	}
	if len(CategoryGoodsList)%pagination.PageSize == 0 {
		pagination.TotalPages = len(CategoryGoodsList) / pagination.PageSize
	} else {
		pagination.TotalPages = len(CategoryGoodsList)/pagination.PageSize + 1
	}

	if pageNum*pagination.PageSize+pagination.PageSize < len(CategoryGoodsList) {
		CategoryGoodsList = CategoryGoodsList[pageNum*pagination.PageSize : pageNum*pagination.PageSize+pagination.PageSize]
	} else {
		CategoryGoodsList = CategoryGoodsList[pageNum*pagination.PageSize:]
	}

	ctx.HTML(http.StatusOK, "shop/category/category.html", gin.H{
		"topNav":            topNav,
		"topNavLen":         len(topNav) - 1,
		"middleNav":         middleNav,
		"footerNav":         footerNav,
		"goodsCate":         goodsCate,
		"focus":             focus,
		"categoryId":        categoryId,
		"categoryGoodsList": CategoryGoodsList,
		"pagination":        pagination,
		"isLogin":           user.UserName != "",
		"user":              user,
	})
}

func (shopController ShopController) Detail(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	goodsId := ctx.Request.FormValue("goodsId")
	goods, goodsImages, goodsAttrs, goodsColor, relationGoods := model.Goods{}, []model.GoodsImage{}, []model.GoodsAttribute{}, []model.GoodsColor{}, []model.Goods{}

	util.Db.Where("id=?", goodsId).Find(&goods)
	util.Db.Where("id IN ?", strings.Split(goods.GoodsImages, " ")).Find(&goodsImages)
	util.Db.Where("goods_id=?", goodsId).Find(&goodsAttrs)
	util.Db.Where("id IN ?", strings.Split(goods.GoodsColor, " ")).Find(&goodsColor)
	util.Db.Where("id IN ?", strings.Split(goods.GoodsRelation, " ")).Find(&relationGoods)
	goodsDescription := template.HTML(goods.GoodsDescription)

	tempColorName := ""
	if len(goodsColor) != 0 {
		tempColorName = goodsColor[0].ColorName
	}

	ctx.HTML(http.StatusOK, "shop/detail/detail.html", gin.H{
		"topNav":           topNav,
		"topNavLen":        len(topNav) - 1,
		"middleNav":        middleNav,
		"footerNav":        footerNav,
		"goodsCate":        goodsCate,
		"focus":            focus,
		"goods":            goods,
		"goodsImages":      goodsImages,
		"goodsColorLen":    len(goodsColor),
		"relationGoods":    relationGoods,
		"goodsColor":       goodsColor,
		"goodsAttrs":       goodsAttrs,
		"goodsDescription": goodsDescription,
		"isLogin":          user.UserName != "",
		"user":             user,
		"tempColorName":    tempColorName,
	})
}

func (shopController ShopController) GetImageList(ctx *gin.Context) {
	goodsImages, colorId, goodsId := []model.GoodsImage{}, ctx.Request.FormValue("colorId"), ctx.Request.FormValue("goodsId")
	util.Db.Where("goods_id=? AND color_id=?", goodsId, colorId).Find(&goodsImages)
	if len(goodsImages) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"imageList": goodsImages,
		})
	}
}

func getPublicInfo(ctx *gin.Context) (topNav []model.Nav, middleNav []model.Nav, footerNav []model.Nav, goodsCate []model.GoodsCate, focus []model.Focus, user model.User) {
	if topNavStr, err := util.Redis.GetRedisStringValue("topNav"); err != nil && topNavStr != "" {
		json.Unmarshal([]byte(topNavStr), &topNav)
	} else {
		util.Db.Where("position=1").Find(&topNav)
		topNavStrTemp, err := json.Marshal(topNav)
		if err != nil {
			fmt.Println("topNavStr Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("topNav", string(topNavStrTemp), 60)
		}
	}

	if tempNavStr, err := util.Redis.GetRedisStringValue("middleNav"); err != nil && tempNavStr != "" {
		json.Unmarshal([]byte(tempNavStr), &middleNav)
	} else {
		util.Db.Where("position=2").Find(&middleNav)
		for i := 0; i < len(middleNav); i++ {
			tempListStr := strings.Split(middleNav[i].Relation, " ")
			util.Db.Where("id in ?", tempListStr).Find(&middleNav[i].RelationGoodsList)
		}
		temp, err := json.Marshal(middleNav)
		if err != nil {
			fmt.Println("middleNav Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("middleNav", string(temp), 60)
		}
	}

	if tempNavStr, err := util.Redis.GetRedisStringValue("footerNav"); err != nil && tempNavStr != "" {
		json.Unmarshal([]byte(tempNavStr), &footerNav)
	} else {
		util.Db.Where("position=3").Find(&footerNav)
		temp, err := json.Marshal(footerNav)
		if err != nil {
			fmt.Println("footerNav Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("footerNav", string(temp), 60)
		}
	}

	if tempNavStr, err := util.Redis.GetRedisStringValue("goodsCate"); err != nil && tempNavStr != "" {
		json.Unmarshal([]byte(tempNavStr), &goodsCate)
	} else {
		util.Db.Preload("GoodsCateList").Where("pid=0").Find(&goodsCate)
		temp, err := json.Marshal(goodsCate)
		if err != nil {
			fmt.Println("goodsCate Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("goodsCate", string(temp), 60)
		}
	}

	if tempNavStr, err := util.Redis.GetRedisStringValue("focus"); err != nil && tempNavStr != "" {
		json.Unmarshal([]byte(tempNavStr), &focus)
	} else {
		util.Db.Find(&focus)
		temp, err := json.Marshal(focus)
		if err != nil {
			fmt.Println("focus Marshal Error")
		} else {
			util.Redis.SetRedisStringValue("focus", string(temp), 60)
		}
	}

	util.GetCookie(ctx, "userinfo", &user)
	return
}
