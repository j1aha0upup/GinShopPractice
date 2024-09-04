package controller

import (
	"fmt"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (goodsTypeAttributeController GoodsTypeAttributeController) GoodsTypeAttributeAdd(ctx *gin.Context) {
	goodsType := []model.GoodsType{}
	goodsTypeId, _ := strconv.Atoi(ctx.Request.FormValue("goodsTypeId"))
	util.Db.Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/goodsTypeAttributeAdd.html", gin.H{
		"goodsTypeId": goodsTypeId,
		"goodsType":   goodsType,
	})
}

func (goodsTypeAttributeController GoodsTypeAttributeController) DoGoodsTypeAttributeAdd(ctx *gin.Context) {
	goodsTypeAttribute := &model.GoodsTypeAttribute{}
	failUrl := "/admin/goodsTypeAttributeAdd"
	successUrl := "/admin/goodsTypeAttributeList?goodsTypeId=" + ctx.Request.FormValue("goodsTypeId")

	goodsTypeAttribute.Title = strings.Trim(ctx.PostForm("goodsTypeAttributeTitle"), " ")
	goodsTypeAttribute.AttributeValue = ctx.PostForm("goodsTypeAttributeValue")
	goodsTypeAttribute.AttributeType, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeType"))
	goodsTypeAttribute.TypeId, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeTypeId"))
	goodsTypeAttribute.Status, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeStatus"))
	goodsTypeAttribute.Sort, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeSort"))
	goodsTypeAttribute.AddTime = int(time.Now().Unix())

	if util.Db.Create(&goodsTypeAttribute).Error != nil {
		goodsTypeAttributeController.Fail(ctx, "商品属性未成功添加", failUrl)
	} else {
		goodsTypeAttributeController.Success(ctx, "商品属性添加成功", successUrl)
	}
}

func (goodsTypeAttributeController GoodsTypeAttributeController) GoodsTypeAttributeList(ctx *gin.Context) {
	goodsTypeId := ctx.Request.FormValue("goodsTypeId")
	goodsTypeAttribute := []model.GoodsTypeAttribute{}
	util.Db.Where("type_id=?", goodsTypeId).Find(&goodsTypeAttribute)
	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/goodsTypeAttributeList.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeId":        goodsTypeId,
	})
}

func (goodsTypeAttributeController GoodsTypeAttributeController) GoodsTypeAttributeEdit(ctx *gin.Context) {
	goodsTypeAttribute := model.GoodsTypeAttribute{}
	goodsType := []model.GoodsType{}
	goodsTypeAttributeId := ctx.Request.FormValue("goodsTypeAttributeId")

	util.Db.Where("id=?", goodsTypeAttributeId).Find(&goodsTypeAttribute)
	util.Db.Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goodsTypeAttribute/goodsTypeAttributeEdit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsType":          goodsType,
	})
}

func (goodsTypeAttributeController GoodsTypeAttributeController) DoGoodsTypeAttributeEdit(ctx *gin.Context) {
	goodsTypeAttribute := &model.GoodsTypeAttribute{}
	goodsTypeAttribute.Id, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeId"))
	util.Db.Where("id=?", goodsTypeAttribute.Id).Find(&goodsTypeAttribute)

	goodsTypeAttribute.Title = strings.Trim(ctx.PostForm("goodsTypeAttributeTitle"), " ")
	goodsTypeAttribute.AttributeValue = ctx.PostForm("goodsTypeAttributeValue")
	goodsTypeAttribute.AttributeType, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeType"))
	goodsTypeAttribute.TypeId, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeTypeId"))
	goodsTypeAttribute.Status, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeStatus"))
	goodsTypeAttribute.Sort, _ = strconv.Atoi(ctx.PostForm("goodsTypeAttributeSort"))

	failUrl := fmt.Sprintf("/admin/goodsTypeAttributeEdit?goodsTypeAttributeId=%d", goodsTypeAttribute.Id)
	successUrl := "/admin/goodsTypeAttributeList"

	if goodsTypeAttribute.Title == "" {
		goodsTypeAttributeController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		if util.Db.Save(&goodsTypeAttribute).Error != nil {
			goodsTypeAttributeController.Fail(ctx, "属性修改失败", failUrl)
		} else {
			goodsTypeAttributeController.Success(ctx, "属性修改成功", successUrl)
		}
	}
}

func (goodsTypeAttributeController GoodsTypeAttributeController) GoodsTypeAttributeDelete(ctx *gin.Context) {
	goodsTypeAttribute := model.GoodsTypeAttribute{}
	goodsTypeAttribute.Id, _ = strconv.Atoi(ctx.Request.FormValue("goodsTypeAttributeId"))
	util.Db.Where("id=?", goodsTypeAttribute.Id).Find(&goodsTypeAttribute)
	failUrl := "/admin/goodsTypeAttributeList"
	successUrl := "/admin/goodsTypeAttributeList"

	if util.Db.Where("id=?", goodsTypeAttribute.Id).Delete(&goodsTypeAttribute).Error != nil {
		goodsTypeAttributeController.Fail(ctx, "删除属性失败", failUrl)
	} else {
		goodsTypeAttributeController.Success(ctx, "删除属性成功", successUrl)
	}
}
