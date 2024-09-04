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

type AccessController struct {
	BaseController
}

func (accessController AccessController) AddAccess(ctx *gin.Context) {
	moduleIdSelection := &[]model.Access{}
	util.Db.Where("module_id=?", 0).Find(&moduleIdSelection)

	ctx.HTML(http.StatusOK, "admin/admin_user/accessAdd.html", gin.H{
		"moduleIdSelection": moduleIdSelection,
	})
}

func (accessController AccessController) DoAccessAdd(ctx *gin.Context) {
	access := &model.Access{}
	failUrl := "/admin/accessAdd"
	successUrl := "/admin/accessList"

	access.ModuleName = strings.Trim(ctx.PostForm("module_name"), " ")
	access.Type, _ = strconv.Atoi(ctx.PostForm("operate_type"))
	if access.Type == 1 {
		access.ActionName = "模块"
	} else if access.Type == 2 {
		access.ActionName = "菜单"
	} else {
		access.ActionName = "动作"
	}
	access.Url = strings.Trim(ctx.PostForm("operate_url"), " ")
	access.ModuleId, _ = strconv.Atoi(ctx.PostForm("module_id"))
	access.Description = strings.Trim(ctx.PostForm("role_description"), " ")
	access.Sort = 100
	access.Status = 1
	access.AddTime = int(time.Now().Unix())

	if access.ModuleName == "" {
		accessController.Fail(ctx, "模块名称不可为空", failUrl)
	} else {
		if util.Db.Create(&access).Error != nil {
			accessController.Fail(ctx, "权限创建失败", failUrl)
		} else {
			accessController.Success(ctx, "权限创建成功", successUrl)
		}
	}
}

func (accessController AccessController) AccessList(ctx *gin.Context) {
	accessList := &[]model.Access{}

	util.Db.Where("module_id=?", 0).Preload("AccessList").Find(&accessList)
	ctx.HTML(http.StatusOK, "admin/admin_user/accessList.html", gin.H{
		"accessList": accessList,
	})
}

func (accessController AccessController) AccessEdit(ctx *gin.Context) {
	access := &model.Access{}
	access_Id := ctx.Request.FormValue("access_id")

	moduleIdSelection := &[]model.Access{}
	util.Db.Where("module_id=?", 0).Find(&moduleIdSelection)
	util.Db.Where("id=?", access_Id).Find(&access)

	ctx.HTML(http.StatusOK, "admin/admin_user/accessEdit.html", gin.H{
		"moduleIdSelection": moduleIdSelection,
		"access":            access,
	})
}

func (accessController AccessController) DoAccessEdit(ctx *gin.Context) {
	access := &model.Access{}
	access_id := ctx.PostForm("id")
	util.Db.Where("id=?", access_id).Find(&access)

	access.ModuleName = strings.Trim(ctx.PostForm("module_name"), " ")
	access.Type, _ = strconv.Atoi(ctx.PostForm("operate_type"))
	if access.Type == 1 {
		access.ActionName = "模块"
	} else if access.Type == 2 {
		access.ActionName = "菜单"
	} else {
		access.ActionName = "动作"
	}
	access.Url = strings.Trim(ctx.PostForm("operate_url"), " ")
	access.ModuleId, _ = strconv.Atoi(ctx.PostForm("module_id"))
	access.Description = strings.Trim(ctx.PostForm("role_description"), " ")

	failUrl := fmt.Sprintf("/admin/accessEdit?access_id=%d", access.Id)
	successUrl := "/admin/accessList"
	if access.ModuleName == "" {
		accessController.Fail(ctx, "模块名不可为空", failUrl)
	} else {
		if util.Db.Save(&access).Error != nil {
			accessController.Fail(ctx, "权限编辑失败", failUrl)
		} else {
			accessController.Success(ctx, "权限编辑成功", successUrl)
		}
	}
}

func (accessController AccessController) AccessDelete(ctx *gin.Context) {
	access := &model.Access{}
	access.Id, _ = strconv.Atoi(ctx.Request.FormValue("access_id"))

	util.Db.Where("id=?", access.Id).Preload("AccessList").Find(&access)
	if access.ModuleId == 0 && len(access.AccessList) != 0 {
		accessController.Fail(ctx, "该模块下存在子节点未删除", "/admin/accessList")
	} else {
		if util.Db.Where("id=?", access.Id).Delete(&access).Error != nil {
			accessController.Fail(ctx, "删除权限失败", "/admin/accessList")
		} else {
			accessController.Success(ctx, "删除权限成功", "/admin/accessList")
		}
	}
}

func (accessController AccessController) MyAuth(ctx *gin.Context) {
	accessList := []model.Access{}
	role := &model.Role{}
	role_id := ctx.Request.FormValue("role_id")

	util.Db.Where("module_id=?", 0).Preload("AccessList").Find(&accessList)
	util.Db.Where("id=?", role_id).Preload("AccessList").Find(&role)

	for _, v1 := range role.AccessList {
		for i2, v2 := range accessList {
			if v1.Id == v2.Id {
				accessList[i2].IsChecked = true
			}
			for i3, v3 := range v2.AccessList {
				if v3.Id == v1.Id {
					accessList[i2].AccessList[i3].IsChecked = true
				}
			}
		}
	}

	ctx.HTML(http.StatusOK, "admin/admin_user/auth.html", gin.H{
		"accessList": accessList,
		"role_id":    role_id,
	})
}

func (accessController AccessController) DoAuth(ctx *gin.Context) {
	roleAccess := []model.RoleAccess{}
	role_id, _ := strconv.Atoi(ctx.PostForm("role_id"))
	util.Db.Where("role_id=?", role_id).Find(&roleAccess)
	util.Db.Delete(&roleAccess)

	module_id_arrayString := ctx.PostFormArray("module_id")
	accessIdArray := []int{}
	for _, v := range module_id_arrayString {
		temp, _ := strconv.Atoi(v)
		accessIdArray = append(accessIdArray, temp)
	}
	tempRoleAccessArray := []model.RoleAccess{}
	for _, v := range accessIdArray {
		tempRoleAccess := model.RoleAccess{
			RoleId:   role_id,
			AccessId: v,
		}
		tempRoleAccessArray = append(tempRoleAccessArray, tempRoleAccess)
	}
	err := util.Db.Save(&tempRoleAccessArray).Error
	if err != nil {
		accessController.Fail(ctx, "授权失败", "/admin/roleList")
	} else {
		accessController.Success(ctx, "授权成功", "/admin/roleList")
	}
}
