package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminUserContoller struct {
	BaseController
}

func (adminUserContoller AdminUserContoller) AddAdminUser(ctx *gin.Context) {
	role := []model.Role{}
	util.Db.Find(&role)

	ctx.HTML(http.StatusOK, "admin/admin_user/addAdminUser.html", gin.H{
		"role": role,
	})
}

func (adminUserContoller AdminUserContoller) DoAddAdminUser(ctx *gin.Context) {
	failUrl := "/admin/addAdminUser"

	user := &model.AdminUser{}
	user.UserName = strings.Trim(ctx.PostForm("username"), " ")
	if user.UserName == "" {
		adminUserContoller.Fail(ctx, "用户名不可为空", failUrl)
	} else {
		tempUser := &model.AdminUser{}
		util.Db.Where("user_name = ?", user.UserName).Find(tempUser)
		if tempUser.Id != 0 {
			adminUserContoller.Fail(ctx, "用户已存在", failUrl)
		} else {
			user.Password = ctx.PostForm("password")
			if len(user.Password) < 6 {
				adminUserContoller.Fail(ctx, "密码不可小于6位", failUrl)
			} else {
				md5Password := md5.New()
				md5Password.Write([]byte(user.Password))
				user.Password = hex.EncodeToString(md5Password.Sum(nil))

				user.Mobile = ctx.PostForm("mobile")
				if user.Mobile == "" {
					adminUserContoller.Fail(ctx, "手机号码不可为空", failUrl)
				} else {
					user.Email = ctx.PostForm("email")
					user.RoleId, _ = strconv.Atoi(ctx.PostForm("role_id"))
					user.Status = 1
					user.IsSuper = false
					user.AddTime = int(time.Now().Unix())
					if util.Db.Create(&user).Error != nil {
						adminUserContoller.Fail(ctx, "增加用户失败", failUrl)
					} else {
						adminUserContoller.Success(ctx, "增加用户成功", failUrl)
					}
				}
			}
		}
	}
}

func (adminUserContoller AdminUserContoller) GetAdminUser(ctx *gin.Context) {
	userList := &[]model.AdminUser{}
	util.Db.Preload("ActRole").Find(&userList)
	ctx.HTML(http.StatusOK, "admin/admin_user/adminUserList.html", gin.H{
		"userlist": userList,
	})
}

func (adminUserContoller AdminUserContoller) EditAdminUser(ctx *gin.Context) {
	user := &model.AdminUser{}
	user_id := ctx.Request.FormValue("user_id")
	role := []model.Role{}
	util.Db.Find(&role)

	util.Db.Where("id=?", user_id).Preload("ActRole").Find(&user)

	ctx.HTML(http.StatusOK, "admin/admin_user/adminUserEdit.html", gin.H{
		"user": user,
		"role": role,
	})
}

func (adminUserContoller AdminUserContoller) DoEditAdminUser(ctx *gin.Context) {
	user := &model.AdminUser{}
	user.UserName = strings.Trim(ctx.PostForm("username"), " ")
	user.Id, _ = strconv.Atoi(ctx.PostForm("id"))

	failUrl := fmt.Sprintf("/admin/editAdminUser?user_id=%d", user.Id)
	successUrl := "/admin/getAdminUser"

	if user.UserName == "" {
		adminUserContoller.Fail(ctx, "用户名不可为空", failUrl)
	} else {
		tempUser := &model.AdminUser{}
		util.Db.Where("id = ?", user.Id).Find(&tempUser)
		if user.UserName != tempUser.UserName {
			user.Id = 0
			util.Db.Where("user_name = ?", user.UserName).Find(&user)

			if user.Id != 0 {
				adminUserContoller.Fail(ctx, "用户已存在", failUrl)
			} else {
				tempPassword := ctx.PostForm("password")
				if len(tempPassword) < 6 && tempPassword != "" {
					adminUserContoller.Fail(ctx, "密码不可小于6位", failUrl)
				} else {
					if tempPassword != "" {
						md5Password := md5.New()
						md5Password.Write([]byte(tempPassword))
						user.Password = hex.EncodeToString(md5Password.Sum(nil))
					} else {
						user.Password = tempUser.Password
					}

					user.Mobile = ctx.PostForm("mobile")
					if user.Mobile == "" {
						adminUserContoller.Fail(ctx, "手机号码不可为空", failUrl)
					} else {
						user.Email = ctx.PostForm("email")
						user.RoleId, _ = strconv.Atoi(ctx.PostForm("role_id"))

						if util.Db.Save(&user).Error != nil {
							adminUserContoller.Fail(ctx, "修改用户失败", failUrl)
						} else {
							adminUserContoller.Success(ctx, "修改用户成功", successUrl)
						}
					}
				}
			}
		} else {
			tempPassword := ctx.PostForm("password")
			if len(tempPassword) < 6 && tempPassword != "" {
				adminUserContoller.Fail(ctx, "密码不可小于6位", failUrl)
			} else {
				if tempPassword != "" {
					md5Password := md5.New()
					md5Password.Write([]byte(tempPassword))
					user.Password = hex.EncodeToString(md5Password.Sum(nil))
				} else {
					user.Password = tempUser.Password
				}

				user.Mobile = ctx.PostForm("mobile")
				if user.Mobile == "" {
					adminUserContoller.Fail(ctx, "手机号码不可为空", failUrl)
				} else {
					user.Email = ctx.PostForm("email")
					user.RoleId, _ = strconv.Atoi(ctx.PostForm("role_id"))

					if util.Db.Save(&user).Error != nil {
						adminUserContoller.Fail(ctx, "修改用户失败", failUrl)
					} else {
						adminUserContoller.Success(ctx, "修改用户成功", successUrl)
					}
				}
			}
		}

	}
}

func (adminUserContoller AdminUserContoller) DeleteAdminUser(ctx *gin.Context) {
	user := &model.AdminUser{}
	user.Id, _ = strconv.Atoi(ctx.Request.FormValue("user_id"))

	if util.Db.Where("id=?", user.Id).Delete(&user).Error != nil {
		adminUserContoller.Fail(ctx, "删除用户失败", "/admin/getAdminUser")
	} else {
		adminUserContoller.Success(ctx, "删除用户成功", "/admin/getAdminUser")
	}
}
