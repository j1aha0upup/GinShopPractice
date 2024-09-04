package main

import (
	"practice_gin_store/model"
	"practice_gin_store/router"
	"practice_gin_store/util"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//GIN默认相关配置
	r := gin.Default()

	//模板函数(要在模板加载前配置)
	r.SetFuncMap(template.FuncMap{
		"Unix2Time":             util.Unix2Time,
		"IsColorInGoods":        model.IsColorInGoods,
		"AttributeOptionList":   model.AttributeOptionList,
		"Markdown2Html":         util.Markdown2Html,
		"GoodsPriceMulGoodsNum": util.GoodsPriceMulGoodsNum,
		"TitleStrCut":           util.TitleStrCut,
	})

	//加载模板
	r.LoadHTMLGlob("./template/**/**/*")
	r.Static("/static", "./static")

	//中间件session配置
	store := cookie.NewStore([]byte("secret123"))
	r.Use(sessions.Sessions("mysession", store))

	//路由相关配置
	router.RouterAdminInit(r)
	router.RouterShopInit(r)

	//监听相关参数
	r.Run(":8090")
}
