package middleware

import (
	"encoding/json"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(ctx *gin.Context) {
	pathname := strings.Split(ctx.Request.URL.String(), "?")[0]
	DontUseLogin := map[string]int{
		"/admin/login":        1,
		"/admin/get_captcha":  1,
		"/admin/verify_login": 1,
	}
	session := sessions.Default(ctx)
	userInfo, ok := session.Get("userinfo").(string)
	if ok && userInfo != "" {
		user := model.AdminUser{}
		err := json.Unmarshal([]byte(userInfo), &user)
		if err != nil || user.Id == 0 {
			if _, ok := DontUseLogin[pathname]; !ok {
				ctx.Redirect(302, "/admin/login")
			}
		}
		ctx.Set("user", user)
		//已登录判断该路径用户是否有权限访问
		DontUseAuth := map[string]int{
			"/admin/index":         1,
			"/admin/iframeContent": 1,
			"/admin/logout":        1,
		}

		if _, ok := DontUseAuth[pathname]; !ok && !user.IsSuper {
			accessRoleList := []model.RoleAccess{}
			util.Db.Where("role_id=?", user.RoleId).Find(&accessRoleList)
			accessIdList := []int{}
			for _, v := range accessRoleList {
				accessIdList = append(accessIdList, v.AccessId)
			}

			accessUrlList := []model.Access{}
			util.Db.Where("id IN ? AND type = ?", accessIdList, 1).Preload("AccessList").Find(&accessUrlList)
			flag := false
			for _, v1 := range accessUrlList {
				if v1.Url == pathname {
					flag = true
				}
				for _, v2 := range v1.AccessList {
					if v2.Url == pathname {
						flag = true
						break
					}
				}
				if flag {
					break
				}
			}

			if !flag {
				ctx.String(http.StatusOK, "拒绝访问")
				ctx.Abort()
			}
		}
	} else {
		//未登录检查当前页面是否不需要登录
		if _, ok := DontUseLogin[pathname]; !ok {
			ctx.Redirect(302, "/admin/login")
		}
	}
}

func ShopLogin(ctx *gin.Context) {
	user := model.User{}
	util.GetCookie(ctx, "userinfo", &user)

	if user.Id != 0 {
		ctx.Set("user", user)
	} else {
		ctx.Redirect(302, "/login")
	}
}
