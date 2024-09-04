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

type GoodsTypeController struct {
	BaseController
}

func (goodsTypeController GoodsTypeController) GoodsTypeAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/goodsType/goodsTypeAdd.html", gin.H{})
}

func (goodsTypeController GoodsTypeController) DoGoodsTypeAdd(ctx *gin.Context) {
	goodsType := &model.GoodsType{}
	failUrl := "/admin/goodsTypeAdd"
	successUrl := "/admin/goodsTypeList"

	goodsType.Title = strings.Trim(ctx.PostForm("goodsTypeTitle"), " ")
	goodsType.Description = strings.Trim(ctx.PostForm("goodsTypeDescription"), " ")
	goodsType.Status, _ = strconv.Atoi(ctx.PostForm("goodsTypeStatus"))
	goodsType.AddTime = int(time.Now().Unix())

	if util.Db.Create(&goodsType).Error != nil {
		goodsTypeController.Fail(ctx, "商品类型未成功添加", failUrl)
	} else {
		goodsTypeController.Success(ctx, "商品类型添加成功", successUrl)
	}
}

func (goodsTypeController GoodsTypeController) GoodsTypeList(ctx *gin.Context) {
	goodsType := []model.GoodsType{}
	util.Db.Find(&goodsType)
	ctx.HTML(http.StatusOK, "admin/goodsType/goodsTypeList.html", gin.H{
		"goodsType": goodsType,
	})
}

func (goodsTypeController GoodsTypeController) GoodsTypeEdit(ctx *gin.Context) {
	goodsType := model.GoodsType{}
	goodsTypeId := ctx.Request.FormValue("goodsTypeId")

	util.Db.Where("id=?", goodsTypeId).Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goodsType/goodsTypeEdit.html", gin.H{
		"goodsType": goodsType,
	})
}

func (goodsTypeController GoodsTypeController) DoGoodsTypeEdit(ctx *gin.Context) {
	goodsType := &model.GoodsType{}
	goodsType.Id, _ = strconv.Atoi(ctx.PostForm("goodsTypeId"))
	util.Db.Where("id=?", goodsType.Id).Find(&goodsType)

	goodsType.Title = strings.Trim(ctx.PostForm("goodsTypeTitle"), " ")
	goodsType.Description = strings.Trim(ctx.PostForm("goodsTypeDescription"), " ")
	goodsType.Status, _ = strconv.Atoi(ctx.PostForm("goodsTypeStatus"))

	failUrl := fmt.Sprintf("/admin/goodsCateEdit?goodsCate_id=%d", goodsType.Id)
	successUrl := "/admin/goodsCateList"

	if goodsType.Title == "" {
		goodsTypeController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		if util.Db.Save(&goodsType).Error != nil {
			goodsTypeController.Fail(ctx, "商品类型信息修改失败", failUrl)
		} else {
			goodsTypeController.Success(ctx, "商品类型信息修改成功", successUrl)
		}
	}
}

func (goodsTypeController GoodsTypeController) GoodsTypeDelete(ctx *gin.Context) {
	goodsType := model.GoodsType{}
	goodsTypeAttribute := []model.GoodsTypeAttribute{}
	goodsType.Id, _ = strconv.Atoi(ctx.Request.FormValue("goodsTypeId"))
	util.Db.Where("id=?", goodsType.Id).Find(&goodsType)
	failUrl := "/admin/goodsTypeList"
	successUrl := "/admin/goodsTypeList"

	if util.Db.Where("id=?", goodsType.Id).Delete(&goodsType).Error != nil {
		goodsTypeController.Fail(ctx, "删除商品类型失败", failUrl)
	} else {
		util.Db.Where("type_id = ?", goodsType.Id).Find(&goodsTypeAttribute)
		tempGoodsTypeAttributeId := []int{}
		for _, v := range goodsTypeAttribute {
			tempGoodsTypeAttributeId = append(tempGoodsTypeAttributeId, v.Id)
		}
		util.Db.Where("id IN ?", tempGoodsTypeAttributeId).Delete(&goodsTypeAttribute)
		goodsTypeController.Success(ctx, "删除商品类型成功", successUrl)
	}
}
