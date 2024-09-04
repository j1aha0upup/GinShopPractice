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

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (focusController FocusController) FocusAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/focusAdd.html", gin.H{})
}

func (focusController FocusController) DoFocusAdd(ctx *gin.Context) {
	focus := &model.Focus{}
	failUrl := "/admin/focusAdd"
	successUrl := "/admin/focusList"

	focus.Title = strings.Trim(ctx.PostForm("focus_title"), " ")
	focus.Link = strings.Trim(ctx.PostForm("focus_link"), " ")
	focus.Sort, _ = strconv.Atoi(ctx.PostForm("focus_sort"))
	focus.Status, _ = strconv.Atoi(ctx.PostForm("focus_status"))
	focus.Type, _ = strconv.Atoi(ctx.PostForm("focus_type"))
	focus.AddTime = int(time.Now().Unix())
	file, _ := ctx.FormFile("focus_img")

	if focus.Title == "" || focus.Link == "" {
		focusController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		if len(file.Filename) > 4 && ((file.Filename[len(file.Filename)-4:] == ".jpg") || (file.Filename[len(file.Filename)-4:] == ".png")) {
			fileName := "static/uploaded/uploadedFocusImages/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + file.Filename

			err := ctx.SaveUploadedFile(file, fileName)
			smallImgName := fileName[:len(fileName)-4] + "small.png"

			if err != nil {
				focusController.Fail(ctx, err.Error(), failUrl)
			} else {
				imgSmallOpen, err := imaging.Open(fileName)
				if err != nil {
					fmt.Println("open fail:" + err.Error())
				}
				err = imaging.Save(imaging.Resize(imgSmallOpen, 150, 150, imaging.Lanczos), smallImgName)
				if err != nil {
					fmt.Println("small pic save fail:" + err.Error())
				}
				focus.FocusImg = fileName
				focus.FocusImgSmall = smallImgName
				if util.Db.Create(&focus).Error != nil {
					focusController.Fail(ctx, "轮播图添加失败", failUrl)
				} else {
					focusController.Success(ctx, "轮播图添加成功", successUrl)
				}
			}
		} else {
			focusController.Fail(ctx, "仅支持png与jpg图像", failUrl)
		}
	}
}

func (focusController FocusController) FocusList(ctx *gin.Context) {
	focus := []model.Focus{}
	util.Db.Find(&focus)
	ctx.HTML(http.StatusOK, "admin/focus/focusList.html", gin.H{
		"focus": focus,
	})
}

func (focusController FocusController) FocusEdit(ctx *gin.Context) {
	focus := &model.Focus{}
	focus_id := ctx.Request.FormValue("focus_id")

	util.Db.Where("id=?", focus_id).Find(&focus)

	ctx.HTML(http.StatusOK, "admin/focus/focusEdit.html", gin.H{
		"focus": focus,
	})
}

func (focusController FocusController) DoFocusEdit(ctx *gin.Context) {
	focus := &model.Focus{}
	focus.Id, _ = strconv.Atoi(ctx.PostForm("focus_id"))
	util.Db.Where("id=?", focus.Id).Find(&focus)

	focus.Title = strings.Trim(ctx.PostForm("focus_title"), " ")
	focus.Sort, _ = strconv.Atoi(strings.Trim(ctx.PostForm("focus_sort"), " "))
	focus.Type, _ = strconv.Atoi(strings.Trim(ctx.PostForm("focus_type"), " "))
	focus.Link = strings.Trim(ctx.PostForm("focus_link"), " ")
	failUrl := fmt.Sprintf("/admin/focusEdit?focus_id=%d", focus.Id)
	successUrl := "/admin/focusList"

	if focus.Title == "" {
		focusController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		file, _ := ctx.FormFile("focus_img")
		if file != nil {
			if len(file.Filename) > 4 && ((file.Filename[len(file.Filename)-4:] == ".jpg") || (file.Filename[len(file.Filename)-4:] == ".png")) {
				fileName := "static/uploaded/uploadedFocusImages/" + strconv.Itoa(time.Now().Year()) + "/" + strconv.Itoa(int(time.Now().Month())) + "/" + strconv.Itoa(time.Now().Day()) + "/" + time.Now().Format("20060102150405") + "-" + file.Filename
				err := ctx.SaveUploadedFile(file, fileName)
				smallImgName := fileName[:len(fileName)-4] + "small.png"

				if err != nil {
					focusController.Fail(ctx, err.Error(), failUrl)
				} else {
					imgSmallOpen, err := imaging.Open(fileName)
					if err != nil {
						fmt.Println("open fail:" + err.Error())
					}
					err = imaging.Save(imaging.Resize(imgSmallOpen, 150, 150, imaging.Lanczos), smallImgName)
					if err != nil {
						fmt.Println("small pic save fail:" + err.Error())
					}
					os.Remove(focus.FocusImg)
					os.Remove(focus.FocusImgSmall)

					focus.FocusImg = fileName
					focus.FocusImgSmall = smallImgName
				}
			}
		}
		if util.Db.Save(&focus).Error != nil {
			focusController.Fail(ctx, "轮播图信息修改失败", failUrl)
		} else {
			focusController.Success(ctx, "轮播图信息修改成功", successUrl)
		}
	}
}

func (focusController FocusController) FocusDelete(ctx *gin.Context) {
	focus := &model.Focus{}
	focus.Id, _ = strconv.Atoi(ctx.Request.FormValue("focus_id"))
	util.Db.Where("id=?", focus.Id).Find(&focus)
	err1 := os.Remove(focus.FocusImg)
	err2 := os.Remove(focus.FocusImgSmall)
	if err1 != nil || err2 != nil {
		focusController.Fail(ctx, "删除轮播图文件失败", "/admin/focusList")
		fmt.Println(err1.Error())
		fmt.Println(err2.Error())
	} else {
		if util.Db.Where("id=?", focus.Id).Delete(&focus).Error != nil {
			focusController.Fail(ctx, "删除轮播图失败", "/admin/focusList")
		} else {
			focusController.Success(ctx, "删除轮播图成功", "/admin/focusList")
		}
	}
}
