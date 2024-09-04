package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserContoller struct {
	BaseController
}

func (userContoller UserContoller) Register(ctx *gin.Context) {
	referer := ctx.Request.Referer()
	ctx.HTML(http.StatusOK, "shop/register/register.html", gin.H{
		"referer": referer,
	})
}

func (userContoller UserContoller) DoRegister(ctx *gin.Context) {
	captchaId := ctx.PostForm("captchaId")
	captchaCode := ctx.PostForm("captchaCode")
	referer := ctx.PostForm("referer")

	if model.VerifyCaptcha(captchaId, captchaCode) {
		user := model.User{}
		user.UserName = strings.Trim(ctx.PostForm("username"), " ")
		user.Password = strings.Trim(ctx.PostForm("password"), " ")
		user.Phone = strings.Trim(ctx.PostForm("phone"), " ")
		user.Email = strings.Trim(ctx.PostForm("email"), " ")
		user.Status = 1
		user.AddTime = int(time.Now().Unix())
		user.LastIp = ctx.ClientIP()
		if user.UserName == "" || user.Password == "" || user.Phone == "" || user.Email == "" {
			ctx.String(http.StatusOK, "带*为必填项")
		} else {
			tempUser := []model.User{}
			util.Db.Where("user_name=?", user.UserName).Find(&tempUser)
			if len(tempUser) != 0 {
				ctx.String(http.StatusOK, "用户名已存在")
			} else {
				if len(ctx.PostForm("password")) < 6 || ctx.PostForm("password") != ctx.PostForm("repassword") {
					ctx.String(http.StatusOK, "密码长度 < 6 或 两次密码输入不一致")
				} else {
					md5Password := md5.New()
					md5Password.Write([]byte(user.Password))
					user.Password = hex.EncodeToString(md5Password.Sum(nil))

					util.Db.Create(&user)
					userinfo, _ := json.Marshal(user)
					util.SetCookie(ctx, "userinfo", string(userinfo), 60*60*24)
					if referer != "" {
						ctx.Redirect(302, referer)
					} else {
						ctx.Redirect(302, "/index")
					}
				}
			}
		}

	} else {
		ctx.String(http.StatusOK, "验证码错误")
	}
}

func (userContoller UserContoller) Login(ctx *gin.Context) {
	referer := ctx.Request.Referer()
	ctx.HTML(http.StatusOK, "shop/login/login.html", gin.H{
		"referer": referer,
	})
}

func (userContoller UserContoller) DoLogin(ctx *gin.Context) {
	captchaId := ctx.PostForm("captchaId")
	captchaCode := ctx.PostForm("captchaCode")
	referer := ctx.PostForm("referer")
	user := []model.User{}

	if model.VerifyCaptcha(captchaId, captchaCode) {
		username := ctx.PostForm("username")
		tempMd5 := md5.New()
		io.WriteString(tempMd5, ctx.PostForm("password"))
		password := string(hex.EncodeToString(tempMd5.Sum(nil)))
		util.Db.Where("user_name=? AND password=?", username, password).Find(&user)

		if len(user) != 0 {
			userinfo, _ := json.Marshal(user[0])
			util.SetCookie(ctx, "userinfo", string(userinfo), 60*60*24)
			if referer != "" {
				ctx.Redirect(302, referer)
			} else {
				ctx.Redirect(302, "/index")
			}
		} else {
			ctx.String(http.StatusOK, "用户名或密码不正确")
		}
	} else {
		ctx.String(http.StatusOK, "验证码错误")
	}
}
