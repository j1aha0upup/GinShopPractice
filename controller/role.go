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

type RoleController struct {
	BaseController
}

func (roleController RoleController) AddRole(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/admin_user/roleAdd.html", gin.H{})
}

func (roleController RoleController) DoAddRole(ctx *gin.Context) {
	role := &model.Role{}
	failUrl := "/admin/addRole"
	successUrl := "/admin/roleList"

	role.Title = strings.Trim(ctx.PostForm("role_name"), " ")
	role.Description = strings.Trim(ctx.PostForm("role_description"), " ")
	role.Status = 1
	role.AddTime = int(time.Now().Unix())

	if role.Title == "" {
		roleController.Fail(ctx, "角色名称不可为空", failUrl)
	} else {
		if util.Db.Create(&role).Error != nil {
			roleController.Fail(ctx, "角色创建失败", failUrl)
		} else {
			roleController.Success(ctx, "角色创建成功", successUrl)
		}
	}
}

func (roleController RoleController) RoleList(ctx *gin.Context) {
	roleList := &[]model.Role{}
	util.Db.Find(&roleList)
	ctx.HTML(http.StatusOK, "admin/admin_user/roleList.html", gin.H{
		"roleList": roleList,
	})
}

func (roleController RoleController) EditRole(ctx *gin.Context) {
	role := &model.Role{}
	role_id := ctx.Request.FormValue("role_id")

	util.Db.Where("id=?", role_id).Find(&role)

	ctx.HTML(http.StatusOK, "admin/admin_user/roleEdit.html", gin.H{
		"role": role,
	})
}

func (roleController RoleController) DoEditRole(ctx *gin.Context) {
	role := &model.Role{}
	role.Id, _ = strconv.Atoi(ctx.PostForm("id"))
	util.Db.Where("id=?", role.Id).Find(&role)

	role.Title = strings.Trim(ctx.PostForm("role_name"), " ")
	role.Description = strings.Trim(ctx.PostForm("description"), " ")
	failUrl := fmt.Sprintf("/admin/editRole?role_id=%d", role.Id)
	successUrl := "/admin/roleList"
	if role.Title == "" {
		roleController.Fail(ctx, "角色名称不可为空", failUrl)
	} else {
		if util.Db.Save(&role).Error != nil {
			roleController.Fail(ctx, "角色编辑失败", failUrl)
		} else {
			roleController.Success(ctx, "角色编辑成功", successUrl)
		}
	}
}

func (roleController RoleController) DeleteRole(ctx *gin.Context) {
	role := &model.Role{}
	role.Id, _ = strconv.Atoi(ctx.Request.FormValue("role_id"))

	if util.Db.Where("id=?", role.Id).Delete(&role).Error != nil {
		roleController.Fail(ctx, "删除角色失败", "/admin/roleList")
	} else {
		roleController.Success(ctx, "删除角色成功", "/admin/roleList")
	}
}
