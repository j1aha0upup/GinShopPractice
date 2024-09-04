package controller

import (
	"fmt"
	"net/http"
	"os"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GoodsCateController struct {
	BaseController
}

func (goodsCateController GoodsCateController) GoodsCateAdd(ctx *gin.Context) {
	goodsCate := []model.GoodsCate{}
	util.Db.Where("pid=0").Find(&goodsCate)
	ctx.HTML(http.StatusOK, "admin/goodsCate/goodsCateAdd.html", gin.H{
		"goodsCate": goodsCate,
	})
}

func (goodsCateController GoodsCateController) DoGoodsCateAdd(ctx *gin.Context) {
	goodsCate := &model.GoodsCate{}
	failUrl := "/admin/goodsCateAdd"
	successUrl := "/admin/goodsCateList"

	goodsCate.Title = strings.Trim(ctx.PostForm("goodsCateTitle"), " ")
	goodsCate.SubTitle = strings.Trim(ctx.PostForm("goodsCateSubTitle"), " ")
	goodsCate.Keywords = strings.Trim(ctx.PostForm("goodsCateKeywords"), " ")
	goodsCate.Link = strings.Trim(ctx.PostForm("goodsCateLink"), " ")
	goodsCate.Description = strings.Trim(ctx.PostForm("goodsCateDescription"), " ")
	goodsCate.Sort, _ = strconv.Atoi(ctx.PostForm("goodsCateSort"))
	goodsCate.Status, _ = strconv.Atoi(ctx.PostForm("goodsCateStatus"))
	goodsCate.Pid, _ = strconv.Atoi(ctx.PostForm("goodsCatePid"))
	goodsCate.AddTime = int(time.Now().Unix())

	file, _ := ctx.FormFile("goodsCateImage")

	if goodsCate.Title == "" || goodsCate.Link == "" {
		goodsCateController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		if len(file.Filename) > 4 && ((file.Filename[len(file.Filename)-4:] == ".jpg") || (file.Filename[len(file.Filename)-4:] == ".png")) {
			fileName := "static/uploaded/uploadedGoodsCate/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + file.Filename

			err := ctx.SaveUploadedFile(file, fileName)
			if err != nil {
				goodsCateController.Fail(ctx, err.Error(), failUrl)
			} else {
				goodsCate.CateImage = fileName

				if util.Db.Create(&goodsCate).Error != nil {
					goodsCateController.Fail(ctx, "商品分类未成功添加", failUrl)
				} else {
					goodsCateController.Success(ctx, "商品分类添加成功", successUrl)
				}
			}
		} else {
			goodsCateController.Fail(ctx, "仅支持png与jpg图像", failUrl)
		}

	}
}

func (goodsCateController GoodsCateController) GoodsCateList(ctx *gin.Context) {
	goodsCate := []model.GoodsCate{}
	util.Db.Where("pid=0").Preload("GoodsCateList").Find(&goodsCate)
	ctx.HTML(http.StatusOK, "admin/goodsCate/goodsCateList.html", gin.H{
		"goodsCate": goodsCate,
	})
}

func (goodsCateController GoodsCateController) GoodsCateEdit(ctx *gin.Context) {
	goodsCate := model.GoodsCate{}
	goodsCateList := []model.GoodsCate{}
	goodsCate_id := ctx.Request.FormValue("goodsCate_id")

	util.Db.Where("id=?", goodsCate_id).Find(&goodsCate)
	util.Db.Where("pid=0").Find(&goodsCateList)

	ctx.HTML(http.StatusOK, "admin/goodsCate/goodsCateEdit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (goodsCateController GoodsCateController) DoGoodsCateEdit(ctx *gin.Context) {
	goodsCate := &model.GoodsCate{}
	goodsCate.Id, _ = strconv.Atoi(ctx.PostForm("goodsCateId"))
	util.Db.Where("id=?", goodsCate.Id).Find(&goodsCate)

	goodsCate.Title = strings.Trim(ctx.PostForm("goodsCateTitle"), " ")
	goodsCate.SubTitle = strings.Trim(ctx.PostForm("goodsCateSubTitle"), " ")
	goodsCate.Keywords = strings.Trim(ctx.PostForm("goodsCateKeywords"), " ")
	goodsCate.Link = strings.Trim(ctx.PostForm("goodsCateLink"), " ")
	goodsCate.Description = strings.Trim(ctx.PostForm("goodsCateDescription"), " ")

	goodsCate.Pid, _ = strconv.Atoi(strings.Trim(ctx.PostForm("goodsCatePid"), " "))
	goodsCate.Sort, _ = strconv.Atoi(strings.Trim(ctx.PostForm("goodsCateSort"), " "))
	goodsCate.Status, _ = strconv.Atoi(strings.Trim(ctx.PostForm("goodsCateStatus"), " "))

	failUrl := fmt.Sprintf("/admin/goodsCateEdit?goodsCate_id=%d", goodsCate.Id)
	successUrl := "/admin/goodsCateList"

	if goodsCate.Title == "" {
		goodsCateController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		file, _ := ctx.FormFile("goodsCateImage")
		if file != nil {
			if len(file.Filename) > 4 && ((file.Filename[len(file.Filename)-4:] == ".jpg") || (file.Filename[len(file.Filename)-4:] == ".png")) {
				fileName := "static/uploaded/uploadedGoodsCate/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + file.Filename
				err := ctx.SaveUploadedFile(file, fileName)

				if err != nil {
					goodsCateController.Fail(ctx, err.Error(), failUrl)
				} else {
					os.Remove(goodsCate.CateImage)

					goodsCate.CateImage = fileName
				}
			}
		}
		if util.Db.Save(&goodsCate).Error != nil {
			goodsCateController.Fail(ctx, "商品分类信息修改失败", failUrl)
		} else {
			goodsCateController.Success(ctx, "商品分类信息修改成功", successUrl)
		}
	}
}

func (goodsCateController GoodsCateController) GoodsCateDelete(ctx *gin.Context) {
	goodsCate := model.GoodsCate{}
	goodsCate.Id, _ = strconv.Atoi(ctx.Request.FormValue("goodsCate_id"))
	util.Db.Where("id=?", goodsCate.Id).Preload("GoodsCateList").Find(&goodsCate)
	failUrl := "/admin/goodsCateList"
	successUrl := "/admin/goodsCateList"

	if len(goodsCate.GoodsCateList) != 0 {
		goodsCateController.Fail(ctx, "该分类下存在未删除子分类", failUrl)
	} else {
		err1 := os.Remove(goodsCate.CateImage)

		if err1 != nil {
			goodsCateController.Fail(ctx, "删除分类图标失败", failUrl)
		} else {
			if util.Db.Where("id=?", goodsCate.Id).Delete(&goodsCate).Error != nil {
				goodsCateController.Fail(ctx, "删除分类失败", failUrl)
			} else {
				goodsCateController.Success(ctx, "删除分类成功", successUrl)
			}
		}
	}
}
