package controller

import (
	"net/http"
	"os"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (goodsController GoodsController) GetGoodsTypeAttribute(ctx *gin.Context) {
	goodsTypeAttribute := []model.GoodsTypeAttribute{}
	goodsTypeId := ctx.Request.FormValue("goodsTypeId")
	util.Db.Where("type_id=?", goodsTypeId).Find(&goodsTypeAttribute)
	ctx.JSON(http.StatusOK, goodsTypeAttribute)
}

func (goodsController GoodsController) GoodsDescriptionImagesUpload(ctx *gin.Context) {
	temp, _ := ctx.FormFile("file")
	tempFileName := "static/uploaded/uploadedGoodsDescription/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + temp.Filename
	err := ctx.SaveUploadedFile(temp, tempFileName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "SaveUploadedFile fail"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SaveUploadedFile success",
			"link":    "/" + tempFileName,
		})
	}
}

func (goodsController GoodsController) GoodsImagesUpload(ctx *gin.Context) {
	temp, _ := ctx.FormFile("file")
	tempFileName := "static/uploaded/uploadedGoodsImages/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + temp.Filename
	err := ctx.SaveUploadedFile(temp, tempFileName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Save UploadedFile fail"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "SaveUploadedFile success",
			"imageSrc": tempFileName,
		})
	}
}

func (goodsController GoodsController) GoodsAdd(ctx *gin.Context) {
	goodsColor := []model.GoodsColor{}
	goodsCate := []model.GoodsCate{}
	goodsType := []model.GoodsType{}

	util.Db.Find(&goodsColor)
	util.Db.Find(&goodsCate)
	util.Db.Find(&goodsType)

	ctx.HTML(http.StatusOK, "admin/goods/goodsAdd.html", gin.H{
		"goodsColor": goodsColor,
		"goodsCate":  goodsCate,
		"goodsType":  goodsType,
	})
}

func (goodsController GoodsController) DoGoodsAdd(ctx *gin.Context) {
	goods := model.Goods{}
	goodsTypeAttribute := []model.GoodsTypeAttribute{}
	goodsAttribute := []model.GoodsAttribute{}
	failUrl := "/admin/goodsAdd"
	successUrl := "/admin/goodsList"

	goods.Title = strings.Trim(ctx.PostForm("goodsTitle"), " ")
	goods.SubTitle = strings.Trim(ctx.PostForm("goodsSubTitle"), " ")
	goods.GoodsSn = strings.Trim(ctx.PostForm("goodsSn"), " ")
	goods.CateId, _ = strconv.Atoi(ctx.PostForm("goodsCateId"))
	goods.ShopPrice, _ = strconv.ParseFloat(ctx.PostForm("goodsShopPrice"), 64)
	goods.MarketPrice, _ = strconv.ParseFloat(ctx.PostForm("goodsMarketPrice"), 64)
	goods.GoodsKeywords = strings.Trim(ctx.PostForm("goodsKeywords"), " ")
	goods.GoodsRelation = strings.Trim(ctx.PostForm("goodsRelation"), " ")
	goods.GoodsFitting = strings.Trim(ctx.PostForm("goodsFitting"), " ")
	goods.GoodsGift = strings.Trim(ctx.PostForm("goodsGift"), " ")
	goods.GoodsVersion = strings.Trim(ctx.PostForm("goodsVersion"), " ")
	goods.MoreChoice = strings.Trim(ctx.PostForm("goodsMoreChoice"), " ")
	goods.GoodsColor = strings.Join(ctx.PostFormArray("goodsColor"), " ")
	goods.GoodsDescription = ctx.PostForm("goodsDescription")
	goodsImageForm, _ := ctx.FormFile("goodsImage")

	if ctx.PostForm("goodsIsNew") != "" {
		goods.IsNew = 1
	}
	if ctx.PostForm("goodsIsHot") != "" {
		goods.IsHot = 1
	}
	if ctx.PostForm("goodsIsBest") != "" {
		goods.IsBest = 1
	}
	goods.Status, _ = strconv.Atoi(ctx.PostForm("goodsStatus"))
	goods.Sort, _ = strconv.Atoi(ctx.PostForm("goodsSort"))
	goods.AddTime = int(time.Now().Unix())

	if util.Db.Create(&goods).Error != nil {
		goodsController.Fail(ctx, "商品基本信息有误", failUrl)
	} else {
		if goodsImageForm != nil && len(goodsImageForm.Filename) > 4 && ((goodsImageForm.Filename[len(goodsImageForm.Filename)-4:] == ".jpg") || (goodsImageForm.Filename[len(goodsImageForm.Filename)-4:] == ".png")) {
			fileName := "static/uploaded/uploadedGoodsImages/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + goodsImageForm.Filename
			ctx.SaveUploadedFile(goodsImageForm, fileName)
			goods.GoodsImage = fileName

			goodsImage := []model.GoodsImage{}
			goodsImageList := ctx.PostFormArray("goodsImageList")
			for _, v := range goodsImageList {
				temp := model.GoodsImage{}
				temp.ImageUrl = v
				temp.GoodsId = goods.Id
				temp.Sort = 100
				temp.Status = 1
				temp.AddTime = int(time.Now().Unix())
				goodsImage = append(goodsImage, temp)
			}
			if len(goodsImage) != 0 && util.Db.Create(&goodsImage).Error != nil {
				goodsController.Fail(ctx, "商品相册有误", failUrl)
			} else {
				tempGoodsImage := []string{}
				for _, v := range goodsImage {
					tempGoodsImageId := strconv.Itoa(v.Id)
					tempGoodsImage = append(tempGoodsImage, tempGoodsImageId)
				}
				goods.GoodsImages = strings.Join(tempGoodsImage, " ")

				goods.GoodsTypeId, _ = strconv.Atoi(ctx.PostForm("goodsTypeId"))
				util.Db.Where("type_id=?", goods.GoodsTypeId).Find(&goodsTypeAttribute)
				for _, v := range goodsTypeAttribute {
					tempGoodsAttribute := model.GoodsAttribute{}
					tempGoodsAttribute.GoodsId = goods.Id
					tempGoodsAttribute.TypeId = goods.GoodsTypeId
					tempGoodsAttribute.AttributeId = v.Id
					tempGoodsAttribute.AttributeType = v.AttributeType
					tempGoodsAttribute.AttributeTitle = v.Title
					tempGoodsAttribute.AttributeValue = ctx.PostForm("goodsAttributeId" + strconv.Itoa(v.Id))
					tempGoodsAttribute.AddTime = int(time.Now().Unix())
					tempGoodsAttribute.Status = v.Status
					tempGoodsAttribute.Sort = v.Sort
					goodsAttribute = append(goodsAttribute, tempGoodsAttribute)
				}
				if len(goodsAttribute) != 0 && util.Db.Create(&goodsAttribute).Error != nil {
					goodsController.Fail(ctx, "商品属性有误", failUrl)
				} else {
					tempGoodsAttrsId := []string{}
					for _, v := range goodsAttribute {
						tempGoodsAttribute := strconv.Itoa(v.Id)
						tempGoodsAttrsId = append(tempGoodsAttrsId, tempGoodsAttribute)
					}
					goods.GoodsAttrs = strings.Join(tempGoodsAttrsId, " ")
					if util.Db.Save(&goods).Error != nil {
						goodsController.Fail(ctx, "商品未能成功添加", failUrl)
					} else {
						goodsController.Success(ctx, "商品添加成功", successUrl)
					}

				}
			}
		} else {
			goodsController.Fail(ctx, "仅支持png与jpg图像", failUrl)
		}
	}
}

func (goodsController GoodsController) GoodsList(ctx *gin.Context) {
	pagination := model.Pagination{}
	goods := []model.Goods{}

	pagination.CurrentPage, _ = strconv.Atoi(ctx.Request.FormValue("pageNumber"))
	if pagination.CurrentPage != 0 {
		pagination.CurrentPage--
	}

	pagination.VisiblePages = 10
	pagination.PageSize = 5
	util.Db.Find(&goods)
	if len(goods)%pagination.PageSize == 0 {
		pagination.TotalPages = len(goods) / pagination.PageSize
	} else {
		pagination.TotalPages = len(goods)/pagination.PageSize + 1
	}
	util.Db.Limit(pagination.PageSize).Offset(pagination.CurrentPage * pagination.PageSize).Find(&goods)
	pagination.CurrentPage++
	ctx.HTML(http.StatusOK, "admin/goods/goodsList.html", gin.H{
		"goods":      goods,
		"pagination": pagination,
	})
}

func (goodsController GoodsController) GoodsEdit(ctx *gin.Context) {
	goods := model.Goods{}
	goodsColor := []model.GoodsColor{}
	goodsCate := []model.GoodsCate{}
	goodsType := []model.GoodsType{}
	goodsAttribute := []model.GoodsAttribute{}
	goodsImages := []model.GoodsImage{}
	goodsColorList := []model.GoodsColor{}

	goodsId := ctx.Request.FormValue("goodsId")
	pageNumber := ctx.Request.FormValue("pageNumber")

	util.Db.Find(&goodsColor)
	util.Db.Find(&goodsCate)
	util.Db.Find(&goodsType)
	util.Db.Where("id=?", goodsId).Find(&goods)
	util.Db.Where("goods_id=?", goods.Id).Find(&goodsAttribute)
	util.Db.Where("id IN ?", strings.Split(goods.GoodsImages, " ")).Find(&goodsImages)
	util.Db.Where("id IN ?", strings.Split(goods.GoodsColor, " ")).Find(&goodsColorList)

	ctx.HTML(http.StatusOK, "admin/goods/goodsEdit.html", gin.H{
		"goods":          goods,
		"goodsColor":     goodsColor,
		"goodsCate":      goodsCate,
		"goodsType":      goodsType,
		"goodsImages":    goodsImages,
		"goodsAttribute": goodsAttribute,
		"goodsColorList": goodsColorList,
		"pageNumber":     pageNumber,
	})
}

func (goodsController GoodsController) DoGoodsEdit(ctx *gin.Context) {
	goods := model.Goods{}
	goodsTypeAttribute := []model.GoodsTypeAttribute{}
	goodsAttribute := []model.GoodsAttribute{}
	goodsId := ctx.Request.FormValue("goodsId")
	pageNumber := ctx.Request.FormValue("pageNumber")

	failUrl := "/admin/goodsEdit?goodsId=" + goodsId
	successUrl := "/admin/goodsList?pageNumber=" + pageNumber
	util.Db.Where("id=?", goodsId).Find(&goods)

	goods.Title = strings.Trim(ctx.PostForm("goodsTitle"), " ")
	goods.SubTitle = strings.Trim(ctx.PostForm("goodsSubTitle"), " ")
	goods.GoodsSn = strings.Trim(ctx.PostForm("goodsSn"), " ")
	goods.CateId, _ = strconv.Atoi(ctx.PostForm("goodsCateId"))
	goods.ShopPrice, _ = strconv.ParseFloat(ctx.PostForm("goodsShopPrice"), 64)
	goods.MarketPrice, _ = strconv.ParseFloat(ctx.PostForm("goodsMarketPrice"), 64)
	goods.GoodsKeywords = strings.Trim(ctx.PostForm("goodsKeywords"), " ")
	goods.GoodsRelation = strings.Trim(ctx.PostForm("goodsRelation"), " ")
	goods.GoodsFitting = strings.Trim(ctx.PostForm("goodsFitting"), " ")
	goods.GoodsGift = strings.Trim(ctx.PostForm("goodsGift"), " ")
	goods.GoodsVersion = strings.Trim(ctx.PostForm("goodsVersion"), " ")
	goods.MoreChoice = strings.Trim(ctx.PostForm("goodsMoreChoice"), " ")
	goods.GoodsColor = strings.Join(ctx.PostFormArray("goodsColor"), " ")
	goods.GoodsDescription = ctx.PostForm("goodsDescription")
	goodsImageForm, _ := ctx.FormFile("goodsImage")
	if ctx.PostForm("goodsIsNew") != "" {
		goods.IsNew = 1
	}
	if ctx.PostForm("goodsIsHot") != "" {
		goods.IsHot = 1
	}
	if ctx.PostForm("goodsIsBest") != "" {
		goods.IsBest = 1
	}
	goods.Status, _ = strconv.Atoi(ctx.PostForm("goodsStatus"))
	goods.Sort, _ = strconv.Atoi(ctx.PostForm("goodsSort"))

	if util.Db.Save(&goods).Error != nil {
		goodsController.Fail(ctx, "商品基本信息有误", failUrl)
	} else {
		goodsImage := []model.GoodsImage{}
		goodsImageList := ctx.PostFormArray("goodsImageList")
		for _, v := range goodsImageList {
			temp := model.GoodsImage{}
			temp.ImageUrl = v
			temp.GoodsId = goods.Id
			temp.Sort = 100
			temp.Status = 1
			temp.AddTime = int(time.Now().Unix())
			goodsImage = append(goodsImage, temp)
		}
		if len(goodsImage) != 0 && util.Db.Create(&goodsImage).Error != nil {
			goodsController.Fail(ctx, "商品相册有误", failUrl)
		} else {
			tempGoodsImage := []string{}
			for _, v := range goodsImage {
				tempGoodsImageId := strconv.Itoa(v.Id)
				tempGoodsImage = append(tempGoodsImage, tempGoodsImageId)
			}
			goods.GoodsImages += " " + strings.Join(tempGoodsImage, " ")
			tempGoodsAttributeList := []model.GoodsAttribute{}
			tempGoodsAttributeIdList := strings.Split(goods.GoodsAttrs, " ")
			util.Db.Where("id IN ?", tempGoodsAttributeIdList).Find(&tempGoodsAttributeList).Delete(&tempGoodsAttributeList)

			goods.GoodsTypeId, _ = strconv.Atoi(ctx.PostForm("goodsTypeId"))
			util.Db.Where("type_id=?", goods.GoodsTypeId).Find(&goodsTypeAttribute)
			for _, v := range goodsTypeAttribute {
				tempGoodsAttribute := model.GoodsAttribute{}
				tempGoodsAttribute.GoodsId = goods.Id
				tempGoodsAttribute.TypeId = goods.GoodsTypeId
				tempGoodsAttribute.AttributeId = v.Id
				tempGoodsAttribute.AttributeType = v.AttributeType
				tempGoodsAttribute.AttributeTitle = v.Title
				tempGoodsAttribute.AttributeValue = ctx.PostForm("goodsAttributeId" + strconv.Itoa(v.Id))
				tempGoodsAttribute.AddTime = int(time.Now().Unix())
				tempGoodsAttribute.Status = v.Status
				tempGoodsAttribute.Sort = v.Sort
				goodsAttribute = append(goodsAttribute, tempGoodsAttribute)
			}

			if len(goodsAttribute) != 0 && util.Db.Create(&goodsAttribute).Error != nil {
				goodsController.Fail(ctx, "商品属性有误", failUrl)
			} else {
				tempGoodsAttrsId := []string{}
				for _, v := range goodsAttribute {
					tempGoodsAttribute := strconv.Itoa(v.Id)
					tempGoodsAttrsId = append(tempGoodsAttrsId, tempGoodsAttribute)
				}
				goods.GoodsAttrs = strings.Join(tempGoodsAttrsId, " ")
				if goodsImageForm != nil {
					if len(goodsImageForm.Filename) > 4 && ((goodsImageForm.Filename[len(goodsImageForm.Filename)-4:] == ".jpg") || (goodsImageForm.Filename[len(goodsImageForm.Filename)-4:] == ".png")) {
						fileName := "static/uploaded/uploadedGoodsImages/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + goodsImageForm.Filename
						ctx.SaveUploadedFile(goodsImageForm, fileName)
						goods.GoodsImage = fileName
						if util.Db.Save(&goods).Error != nil {
							goodsController.Fail(ctx, "商品未能成功修改", failUrl)
						} else {
							goodsController.Success(ctx, "商品修改成功", successUrl)
						}
					} else {
						goodsController.Fail(ctx, "仅支持png与jpg图像", failUrl)
					}
				} else {
					if util.Db.Save(&goods).Error != nil {
						goodsController.Fail(ctx, "商品未能成功修改", failUrl)
					} else {
						goodsController.Success(ctx, "商品修改成功", successUrl)
					}
				}
			}
		}
	}
}

func (goodsController GoodsController) GoodsDelete(ctx *gin.Context) {
	goods := model.Goods{}
	goods.Id, _ = strconv.Atoi(ctx.Request.FormValue("goodsId"))
	util.Db.Where("id=?", goods.Id).Find(&goods)
	failUrl := "/admin/goodsList"
	successUrl := "/admin/goodsList"

	tempGoodsAttributeList := []model.GoodsAttribute{}
	tempGoodsAttributeIdList := strings.Split(goods.GoodsAttrs, " ")
	util.Db.Where("id IN ?", tempGoodsAttributeIdList).Find(&tempGoodsAttributeList).Delete(&tempGoodsAttributeList)
	goodsImageList := []model.GoodsImage{}
	util.Db.Where("goods_id=?", goods.Id).Find(&goodsImageList).Delete(&goodsImageList)
	if util.Db.Delete(&goods).Error != nil {
		goodsController.Fail(ctx, "删除商品失败", failUrl)
	} else {
		goodsController.Success(ctx, "删除商品成功", successUrl)
	}
}

func (goodsController GoodsController) GoodsImageDelete(ctx *gin.Context) {
	goods := model.Goods{}
	goodsImage := model.GoodsImage{}
	goodsImageId := ctx.Request.FormValue("goodsImageId")
	goodsId := ctx.Request.FormValue("goodsId")
	err1 := util.Db.Where("id=?", goodsImageId).Find(&goodsImage).Error
	err2 := util.Db.Where("id=?", goodsId).Find(&goods).Error
	err3 := os.Remove(goodsImage.ImageUrl)
	if err1 != nil || err2 != nil || err3 != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
	}
	goodsImageList := strings.Split(goods.GoodsImages, " ")
	newGoodsImage := ""
	for _, v := range goodsImageList {
		if v != strconv.Itoa(goodsImage.Id) {
			newGoodsImage += v + " "
		}
	}
	goods.GoodsImages = strings.Trim(newGoodsImage, " ")
	err1 = util.Db.Save(goods).Error
	err2 = util.Db.Delete(&goodsImage).Error
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func (goodsController GoodsController) GoodsColorBindImages(ctx *gin.Context) {
	goodsImageId, goodsColorId := ctx.Request.FormValue("goodsImageId"), ctx.Request.FormValue("goodsColorId")
	goodsImage := model.GoodsImage{}
	util.Db.Where("id=?", goodsImageId).Find(&goodsImage)
	goodsImage.ColorId, _ = strconv.Atoi(goodsColorId)
	util.Db.Save(&goodsImage)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
