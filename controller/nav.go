package controller

import (
	"fmt"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type NavController struct {
	BaseController
}

func (navController NavController) NavAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/nav/navAdd.html", gin.H{})
}

func (navController NavController) DoNavAdd(ctx *gin.Context) {
	nav := model.Nav{}
	failUrl := "/admin/navAdd"
	successUrl := "/admin/navList"

	nav.Title = strings.Trim(ctx.PostForm("navTitle"), " ")
	nav.Position, _ = strconv.Atoi(ctx.PostForm("navPosition"))
	nav.Link = strings.Trim(ctx.PostForm("navLink"), " ")
	nav.Sort, _ = strconv.Atoi(ctx.PostForm("navSort"))
	nav.Relation = strings.Trim(ctx.PostForm("navRelation"), " ")
	nav.IsOpennew, _ = strconv.Atoi(ctx.PostForm("navIsNewOpen"))
	nav.Status, _ = strconv.Atoi(ctx.PostForm("navStatus"))

	if util.Db.Create(&nav).Error != nil {
		navController.Fail(ctx, "导航未成功添加", failUrl)
	} else {
		navController.Success(ctx, "导航添加成功", successUrl)
	}
}

func (navController NavController) NavList(ctx *gin.Context) {
	pagination := model.Pagination{}
	nav := []model.Nav{}
	pagination.CurrentPage, _ = strconv.Atoi(ctx.Request.FormValue("pageNumber"))
	if pagination.CurrentPage != 0 {
		pagination.CurrentPage--
	}
	pagination.VisiblePages = 10
	pagination.PageSize = 5
	util.Db.Find(&nav)
	if len(nav)%pagination.PageSize == 0 {
		pagination.TotalPages = len(nav) / pagination.PageSize
	} else {
		pagination.TotalPages = len(nav)/pagination.PageSize + 1
	}
	util.Db.Limit(pagination.PageSize).Offset(pagination.CurrentPage * pagination.PageSize).Find(&nav)

	pagination.CurrentPage++
	ctx.HTML(http.StatusOK, "admin/nav/navList.html", gin.H{
		"nav":        nav,
		"pagination": pagination,
	})
}

func (navController NavController) NavEdit(ctx *gin.Context) {
	nav := model.Nav{}
	pageNumber := ctx.Request.FormValue("pageNumber")
	navId := ctx.Request.FormValue("navId")

	util.Db.Where("id=?", navId).Find(&nav)

	ctx.HTML(http.StatusOK, "admin/nav/navEdit.html", gin.H{
		"nav":        nav,
		"pageNumber": pageNumber,
	})
}

func (navController NavController) DoNavEdit(ctx *gin.Context) {
	nav := model.Nav{}
	nav.Id, _ = strconv.Atoi(ctx.PostForm("navId"))
	util.Db.Where("id=?", nav.Id).Find(&nav)

	nav.Title = strings.Trim(ctx.PostForm("navTitle"), " ")
	nav.Position, _ = strconv.Atoi(ctx.PostForm("navPosition"))
	nav.Link = strings.Trim(ctx.PostForm("navLink"), " ")
	nav.Sort, _ = strconv.Atoi(ctx.PostForm("navSort"))
	nav.Relation = strings.Trim(ctx.PostForm("navRelation"), " ")
	nav.IsOpennew, _ = strconv.Atoi(ctx.PostForm("navIsNewOpen"))
	nav.Status, _ = strconv.Atoi(ctx.PostForm("navStatus"))
	pageNumber, _ := strconv.Atoi(ctx.PostForm("pageNumber"))

	failUrl := fmt.Sprintf("/admin/navEdit?navId=%d&pageNumber=%d", nav.Id, pageNumber)
	successUrl := "/admin/goodsCateList?pageNumber=" + strconv.Itoa(pageNumber)

	if nav.Title == "" {
		navController.Fail(ctx, "带*的选项不可为空", failUrl)
	} else {
		if util.Db.Save(&nav).Error != nil {
			navController.Fail(ctx, "导航信息修改失败", failUrl)
		} else {
			navController.Success(ctx, "导航信息修改成功", successUrl)
		}
	}
}

func (navController NavController) NavDelete(ctx *gin.Context) {
	nav := model.Nav{}

	nav.Id, _ = strconv.Atoi(ctx.Request.FormValue("navId"))
	pageNumber := ctx.Request.FormValue("pageNumber")
	util.Db.Where("id=?", nav.Id).Find(&nav)
	failUrl := "/admin/goodsTypeList?pageNumber=" + pageNumber
	successUrl := "/admin/goodsTypeList?pageNumber=" + pageNumber

	if util.Db.Where("id=?", nav.Id).Delete(&nav).Error != nil {
		navController.Fail(ctx, "删除导航失败", failUrl)
	} else {
		navController.Success(ctx, "删除导航成功", successUrl)
	}
}
