package router

import (
	"practice_gin_store/controller"
	"practice_gin_store/middleware"

	"github.com/gin-gonic/gin"
)

func RouterAdminInit(router *gin.Engine) {
	Admin := router.Group("/admin")
	Admin.Use(middleware.AuthRequired)
	{
		Admin.GET("/changeStatus", controller.BaseController{}.ChangeStatus)
		Admin.GET("/clearRedis", controller.BaseController{}.ClearRedis)

		Admin.GET("/index", controller.AdminContoller{}.Index)
		Admin.GET("/login", controller.AdminContoller{}.Login)
		Admin.GET("/get_captcha", controller.AdminContoller{}.GetCaptcha)
		Admin.POST("/verify_login", controller.AdminContoller{}.VerifyLogin)
		Admin.GET("/logout", controller.AdminContoller{}.Logout)
		Admin.GET("/iframeContent", controller.AdminContoller{}.IframeContent)

		Admin.GET("/addAdminUser", controller.AdminUserContoller{}.AddAdminUser)
		Admin.POST("/doAddAdminUser", controller.AdminUserContoller{}.DoAddAdminUser)
		Admin.GET("/getAdminUser", controller.AdminUserContoller{}.GetAdminUser)
		Admin.GET("/editAdminUser", controller.AdminUserContoller{}.EditAdminUser)
		Admin.GET("/deleteAdminUser", controller.AdminUserContoller{}.DeleteAdminUser)
		Admin.POST("/doEditAdminUser", controller.AdminUserContoller{}.DoEditAdminUser)

		Admin.GET("/addRole", controller.RoleController{}.AddRole)
		Admin.POST("/doAddRole", controller.RoleController{}.DoAddRole)
		Admin.GET("/roleList", controller.RoleController{}.RoleList)
		Admin.GET("/editRole", controller.RoleController{}.EditRole)
		Admin.POST("/doEditRole", controller.RoleController{}.DoEditRole)
		Admin.GET("/deleteRole", controller.RoleController{}.DeleteRole)

		Admin.GET("/addAccess", controller.AccessController{}.AddAccess)
		Admin.POST("/doAccessAdd", controller.AccessController{}.DoAccessAdd)
		Admin.GET("/accessList", controller.AccessController{}.AccessList)
		Admin.GET("/accessEdit", controller.AccessController{}.AccessEdit)
		Admin.POST("/doAccessEdit", controller.AccessController{}.DoAccessEdit)
		Admin.GET("/accessDelete", controller.AccessController{}.AccessDelete)
		Admin.GET("/myAuth", controller.AccessController{}.MyAuth)
		Admin.POST("/doAuth", controller.AccessController{}.DoAuth)

		Admin.GET("/focusAdd", controller.FocusController{}.FocusAdd)
		Admin.GET("/focusDelete", controller.FocusController{}.FocusDelete)
		Admin.GET("/focusEdit", controller.FocusController{}.FocusEdit)
		Admin.GET("/focusList", controller.FocusController{}.FocusList)
		Admin.POST("/doFocusAdd", controller.FocusController{}.DoFocusAdd)
		Admin.POST("/doFocusEdit", controller.FocusController{}.DoFocusEdit)

		Admin.GET("/goodsCateAdd", controller.GoodsCateController{}.GoodsCateAdd)
		Admin.GET("/goodsCateDelete", controller.GoodsCateController{}.GoodsCateDelete)
		Admin.GET("/goodsCateEdit", controller.GoodsCateController{}.GoodsCateEdit)
		Admin.GET("/goodsCateList", controller.GoodsCateController{}.GoodsCateList)
		Admin.POST("/doGoodsCateAdd", controller.GoodsCateController{}.DoGoodsCateAdd)
		Admin.POST("/doGoodsCateEdit", controller.GoodsCateController{}.DoGoodsCateEdit)

		Admin.GET("/goodsTypeAdd", controller.GoodsTypeController{}.GoodsTypeAdd)
		Admin.GET("/goodsTypeDelete", controller.GoodsTypeController{}.GoodsTypeDelete)
		Admin.GET("/goodsTypeEdit", controller.GoodsTypeController{}.GoodsTypeEdit)
		Admin.GET("/goodsTypeList", controller.GoodsTypeController{}.GoodsTypeList)
		Admin.POST("/doGoodsTypeAdd", controller.GoodsTypeController{}.DoGoodsTypeAdd)
		Admin.POST("/doGoodsTypeEdit", controller.GoodsTypeController{}.DoGoodsTypeEdit)

		Admin.GET("/goodsTypeAttributeAdd", controller.GoodsTypeAttributeController{}.GoodsTypeAttributeAdd)
		Admin.GET("/goodsTypeAttributeDelete", controller.GoodsTypeAttributeController{}.GoodsTypeAttributeDelete)
		Admin.GET("/goodsTypeAttributeEdit", controller.GoodsTypeAttributeController{}.GoodsTypeAttributeEdit)
		Admin.GET("/goodsTypeAttributeList", controller.GoodsTypeAttributeController{}.GoodsTypeAttributeList)
		Admin.POST("/doGoodsTypeAttributeAdd", controller.GoodsTypeAttributeController{}.DoGoodsTypeAttributeAdd)
		Admin.POST("/doGoodsTypeAttributeEdit", controller.GoodsTypeAttributeController{}.DoGoodsTypeAttributeEdit)

		Admin.GET("/getGoodsTypeAttribute", controller.GoodsController{}.GetGoodsTypeAttribute)
		Admin.GET("/goodsAdd", controller.GoodsController{}.GoodsAdd)
		Admin.GET("/goodsDelete", controller.GoodsController{}.GoodsDelete)
		Admin.GET("/goodsEdit", controller.GoodsController{}.GoodsEdit)
		Admin.GET("/goodsList", controller.GoodsController{}.GoodsList)
		Admin.GET("/goodsImageDelete", controller.GoodsController{}.GoodsImageDelete)
		Admin.GET("/goodsColorBindImages", controller.GoodsController{}.GoodsColorBindImages)
		Admin.POST("/doGoodsAdd", controller.GoodsController{}.DoGoodsAdd)
		Admin.POST("/goodsImagesUpload", controller.GoodsController{}.GoodsImagesUpload)
		Admin.POST("/doGoodsEdit", controller.GoodsController{}.DoGoodsEdit)
		Admin.POST("/goodsDescriptionImagesUpload", controller.GoodsController{}.GoodsDescriptionImagesUpload)

		Admin.GET("/navAdd", controller.NavController{}.NavAdd)
		Admin.GET("/navDelete", controller.NavController{}.NavDelete)
		Admin.GET("/navEdit", controller.NavController{}.NavEdit)
		Admin.GET("/navList", controller.NavController{}.NavList)
		Admin.POST("/doNavAdd", controller.NavController{}.DoNavAdd)
		Admin.POST("/doNavEdit", controller.NavController{}.DoNavEdit)
	}
}
