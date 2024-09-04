package router

import (
	"practice_gin_store/controller"
	"practice_gin_store/middleware"

	"github.com/gin-gonic/gin"
)

func RouterShopInit(router *gin.Engine) {
	Shop := router.Group("")
	// Admin.Use(middleware.AuthRequired)
	{
		Shop.GET("/index", controller.ShopController{}.Index)
		Shop.GET("/category:categoryId", controller.ShopController{}.Category)
		Shop.GET("/detail", controller.ShopController{}.Detail)
		Shop.GET("/product/getImageList", controller.ShopController{}.GetImageList)
		Shop.POST("/doAddCart", controller.CartController{}.DoAddCart)
		Shop.GET("/register", controller.UserContoller{}.Register)
		Shop.POST("/doRegister", controller.UserContoller{}.DoRegister)
		Shop.GET("/login", controller.UserContoller{}.Login)
		Shop.POST("/doLogin", controller.UserContoller{}.DoLogin)
		Shop.GET("/cart", controller.CartController{}.CartList)
		Shop.GET("/cartDeleteGoods", controller.CartController{}.CartDeleteGoods)
		Shop.GET("/cart/changeAllCart", controller.CartController{}.ChangeAllCart)
		Shop.GET("/cart/changeOneChecked", controller.CartController{}.ChangeOneChecked)

		Shop.GET("/cart/checkout", middleware.ShopLogin, controller.CartController{}.Checkout)
		Shop.POST("/address/addAddress", middleware.ShopLogin, controller.AddressController{}.AddAddress)
		Shop.POST("/address/editAddress", middleware.ShopLogin, controller.AddressController{}.EditAddress)
		Shop.GET("/address/getOneAddressList", middleware.ShopLogin, controller.AddressController{}.GetOneAddressList)
		Shop.GET("/address/changeDefaultAddress", middleware.ShopLogin, controller.AddressController{}.ChangeDefaultAddress)

		Shop.POST("/buy/doCheckout", middleware.ShopLogin, controller.CartController{}.DoCheckout)
		Shop.GET("/usercenter", middleware.ShopLogin, controller.UserCenterController{}.UserCenter)
		Shop.GET("/user/order", middleware.ShopLogin, controller.UserCenterController{}.UserOrder)
		Shop.GET("/user/orderinfo", middleware.ShopLogin, controller.UserCenterController{}.UserOrderInfo)
	}
}
