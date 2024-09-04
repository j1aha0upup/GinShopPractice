package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminContoller struct {
	BaseController
}

func (adminContoller AdminContoller) Index(ctx *gin.Context) {
	userTemp, _ := ctx.Get("user")
	user, _ := userTemp.(model.AdminUser)
	accessRoleList := []model.RoleAccess{}
	accessUrlList := []model.Access{}
	if user.IsSuper {
		util.Db.Where("type = ?", 1).Preload("AccessList", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessUrlList)
	} else {
		util.Db.Where("role_id=?", user.RoleId).Find(&accessRoleList)
		accessIdList := []int{}
		for _, v := range accessRoleList {
			accessIdList = append(accessIdList, v.AccessId)
		}
		util.Db.Where("id IN ? AND type = ?", accessIdList, 1).Preload("AccessList", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessUrlList)
	}
	ctx.HTML(http.StatusOK, "admin/public/index.html", gin.H{
		"user":          user,
		"accessUrlList": accessUrlList,
	})
}

func (adminContoller AdminContoller) Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/public/login.html", gin.H{})
}

func (adminContoller AdminContoller) GetCaptcha(ctx *gin.Context) {
	newCaptcha := &model.MyCaptcha{}
	newCaptcha.CreateCaptcha()

	ctx.JSON(http.StatusOK, newCaptcha)
}

func (adminContoller AdminContoller) VerifyLogin(ctx *gin.Context) {
	captchaId := ctx.PostForm("captchaId")
	captchaCode := ctx.PostForm("captchaCode")
	username := ctx.PostForm("username")
	//获取到form表单中的password并MD5加密
	tempMd5 := md5.New()
	password := ctx.PostForm("password")
	io.WriteString(tempMd5, password)
	password = string(hex.EncodeToString(tempMd5.Sum(nil)))

	if model.VerifyCaptcha(captchaId, captchaCode) {
		user := []model.AdminUser{}
		util.Db.Where("user_name=? AND password=?", username, password).Preload("ActRole").Find(&user)

		if len(user) != 0 {
			session := sessions.Default(ctx)
			userInfo, _ := json.Marshal(user[0])

			session.Set("userinfo", string(userInfo))
			session.Options(sessions.Options{MaxAge: 60 * 60 * 24})
			session.Save()

			adminContoller.Success(ctx, "登陆成功", "/admin/index")
		} else {
			adminContoller.Fail(ctx, "用户名或密码错误", "/admin/login")
		}
	} else {
		adminContoller.Fail(ctx, "验证码错误", "/admin/login")
	}
}

func (adminContoller AdminContoller) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userInfo := session.Get("userinfo")
	if userInfo != nil {
		session.Delete("userinfo")
		session.Save()
		session.Options(sessions.Options{MaxAge: -1})
		session.Save()
		adminContoller.Success(ctx, "已退出", "/admin/login")
	} else {
		adminContoller.Fail(ctx, "请重新确认退出", "/admin/index")
	}
}

func (adminContoller AdminContoller) IframeContent(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/public/iframeContent.html", gin.H{})
}
