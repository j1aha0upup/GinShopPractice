package controller

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	BaseController
}

func (cartController CartController) DoAddCart(ctx *gin.Context) {
	cart, tempCart := []model.Cart{}, model.Cart{}
	tempCart.GoodsId, _ = strconv.Atoi(strings.Trim(ctx.PostForm("goodsId"), " "))
	tempCart.Price, _ = strconv.ParseFloat(strings.Trim(ctx.PostForm("goodsPrice"), " "), 64)
	tempCart.Title = strings.Trim(ctx.PostForm("goodsTitle"), " ")
	tempCart.GoodsVersion = strings.Trim(ctx.PostForm("goodsVersion"), " ")
	tempCart.GoodsGift = strings.Trim(ctx.PostForm("goodsGift"), " ")
	tempCart.GoodsFitting = strings.Trim(ctx.PostForm("goodsFitting"), " ")
	tempCart.GoodsImage = strings.Trim(ctx.PostForm("goodsImage"), " ")
	tempCart.GoodsAttr = strings.Trim(ctx.PostForm("goodsAttr"), " ")
	tempCart.GoodsColor = strings.Trim(ctx.PostForm("goodsColor"), " ")
	tempCart.Num = 1
	tempCart.Checked = true

	util.GetCookie(ctx, "cart", &cart)
	i := 0

	for ; i < len(cart); i++ {
		if cart[i].GoodsId == tempCart.GoodsId && tempCart.GoodsColor == cart[i].GoodsColor {
			cart[i].Num++
			break
		}
	}

	if i == len(cart) {
		cart = append(cart, tempCart)
	}
	cartStr, _ := json.Marshal(cart)
	util.SetCookie(ctx, "cart", string(cartStr), 60*60*24*7)
	ctx.HTML(http.StatusOK, "shop/cart/addcartSuccess.html", gin.H{
		"tempCart": tempCart,
	})
}

func (cartController CartController) CartList(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	cart := []model.Cart{}
	util.GetCookie(ctx, "cart", &cart)

	goodsCheckedLen, totalPrice := 0, 0.0
	for _, v := range cart {
		if v.Checked {
			goodsCheckedLen++
			totalPrice += v.Price * float64(v.Num)
		}
	}

	ctx.HTML(http.StatusOK, "shop/cart/cart.html", gin.H{
		"topNav":          topNav,
		"topNavLen":       len(topNav) - 1,
		"middleNav":       middleNav,
		"footerNav":       footerNav,
		"goodsCate":       goodsCate,
		"focus":           focus,
		"isLogin":         user.UserName != "",
		"user":            user,
		"cart":            cart,
		"cartLen":         len(cart),
		"goodsCheckedLen": goodsCheckedLen,
		"totalPrice":      totalPrice,
	})
}

func (cartController CartController) CartDeleteGoods(ctx *gin.Context) {
	goodsId, _ := strconv.Atoi(ctx.Request.FormValue("goodsId"))
	colorName := ctx.Request.FormValue("colorName")
	cart, resultCart := []model.Cart{}, []model.Cart{}
	util.GetCookie(ctx, "cart", &cart)

	i := 0
	for ; i < len(cart); i++ {
		if cart[i].GoodsId == goodsId && cart[i].GoodsColor == colorName {
			continue
		} else {
			resultCart = append(resultCart, cart[i])
		}
	}
	resultCartStr, _ := json.Marshal(resultCart)

	util.SetCookie(ctx, "cart", string(resultCartStr), 60*60*24*7)

	ctx.Redirect(302, "/cart")
}

func (cartController CartController) ChangeAllCart(ctx *gin.Context) {
	flag := ctx.Request.FormValue("flag")
	cart := []model.Cart{}
	util.GetCookie(ctx, "cart", &cart)

	if flag != "0" {
		for i := 0; i < len(cart); i++ {
			cart[i].Checked = true
		}
	} else {
		for i := 0; i < len(cart); i++ {
			cart[i].Checked = false
		}
	}

	goodsCheckedLen, totalPrice := 0, 0.0
	for _, v := range cart {
		if v.Checked {
			goodsCheckedLen++
			totalPrice += v.Price * float64(v.Num)
		}
	}

	cartStr, _ := json.Marshal(cart)
	util.SetCookie(ctx, "cart", string(cartStr), 60*60*24*7)
	ctx.JSON(http.StatusOK, gin.H{
		"success":         true,
		"goodsCheckedLen": goodsCheckedLen,
		"totalPrice":      totalPrice,
	})
}

func (cartController CartController) ChangeOneChecked(ctx *gin.Context) {
	flag := ctx.Request.FormValue("flag")
	goodsId, _ := strconv.Atoi(ctx.Request.FormValue("goodsId"))
	goodsColor := ctx.Request.FormValue("goodsColor")
	cart := []model.Cart{}
	util.GetCookie(ctx, "cart", &cart)

	for i := 0; i < len(cart); i++ {
		if cart[i].GoodsColor == goodsColor && goodsId == cart[i].GoodsId {
			if flag == "1" {
				cart[i].Checked = true
			} else {
				cart[i].Checked = false
			}
			break
		}
	}

	goodsCheckedLen, totalPrice := 0, 0.0
	for _, v := range cart {
		if v.Checked {
			goodsCheckedLen++
			totalPrice += v.Price * float64(v.Num)
		}
	}

	cartStr, _ := json.Marshal(cart)
	util.SetCookie(ctx, "cart", string(cartStr), 60*60*24*7)
	ctx.JSON(http.StatusOK, gin.H{
		"success":         true,
		"goodsCheckedLen": goodsCheckedLen,
		"totalPrice":      totalPrice,
	})
}

func (cartController CartController) Checkout(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	cart := []model.Cart{}
	util.GetCookie(ctx, "cart", &cart)
	address := []model.Address{}
	util.Db.Where("uid=?", user.Id).Order("id desc").Find(&address)
	totalNum, totalPrice := 0, 0.0
	for _, v := range cart {
		if v.Checked {
			totalNum++
			totalPrice += v.Price * float64(v.Num)
		}
	}

	if totalNum == 0 {
		ctx.JSON(http.StatusOK, "没有选中的商品")
		return
	}

	hash := md5.New()
	hash.Write([]byte(time.Now().Format("20060102150405")))
	paysign := hex.EncodeToString(hash.Sum(nil))
	sessionClient := sessions.Default(ctx)
	sessionClient.Set("paysign", paysign)
	sessionClient.Save()

	ctx.HTML(http.StatusOK, "shop/checkout/checkout.html", gin.H{
		"topNav":     topNav,
		"topNavLen":  len(topNav) - 1,
		"middleNav":  middleNav,
		"footerNav":  footerNav,
		"goodsCate":  goodsCate,
		"focus":      focus,
		"isLogin":    user.UserName != "",
		"user":       user,
		"cart":       cart,
		"address":    address,
		"totalNum":   totalNum,
		"totalPrice": totalPrice,
		"paysign":    paysign,
	})
}

func (cartController CartController) DoCheckout(ctx *gin.Context) {
	paysignForm := ctx.PostForm("paysign")
	temp := sessions.Default(ctx).Get("paysign")
	paysignSession, _ := temp.(string)
	if paysignForm != paysignSession {
		ctx.Redirect(302, "/cart")
		return
	}

	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	cart, newCart := []model.Cart{}, []model.Cart{}
	order, orderItemList := model.Order{}, []model.OrderItem{}
	util.GetCookie(ctx, "cart", &cart)
	address := model.Address{}
	util.Db.Where("uid=? and default_address=?", user.Id, 1).Find(&address)

	totalPrice := 0.0
	for _, v := range cart {
		if v.Checked {
			orderItem := model.OrderItem{
				AddTime:      int(time.Now().Unix()),
				GoodsId:      v.GoodsId,
				GoodsTitle:   v.Title,
				GoodsImage:   v.GoodsImage,
				GoodsPrice:   v.Price,
				GoodsNum:     v.Num,
				GoodsVersion: v.GoodsVersion,
				GoodsColor:   v.GoodsColor,
				Uid:          user.Id,
			}
			orderItemList = append(orderItemList, orderItem)
			totalPrice += v.Price * float64(v.Num)
		} else {
			newCart = append(newCart, v)
		}
	}
	util.Db.Create(&orderItemList)
	if address.Id == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "adress error",
		})
		return
	} else {
		order.Name = address.Name
		order.OrderSn = time.Now().Format("20060102150405") + strconv.Itoa(user.Id)
		order.AddTime = int(time.Now().Unix())
		order.OrderStatus = 1
		order.PayStatus = 0
		order.PayType = 0
		order.Phone = address.Phone
		order.Address = address.Address
		order.TotalPrice = totalPrice
		order.Uid = user.Id
	}
	util.Db.Create(&order)

	for i := 0; i < len(orderItemList); i++ {
		orderItemList[i].OrderId = order.Id
	}
	util.Db.Save(&orderItemList)

	newCartByte, _ := json.Marshal(newCart)

	util.SetCookie(ctx, "cart", string(newCartByte), 60*60*24*7)
	sessions.Default(ctx).Delete("paysign")
	sessions.Default(ctx).Save()
	ctx.HTML(http.StatusOK, "shop/checkout/pay.html", gin.H{
		"topNav":        topNav,
		"topNavLen":     len(topNav) - 1,
		"middleNav":     middleNav,
		"footerNav":     footerNav,
		"goodsCate":     goodsCate,
		"focus":         focus,
		"isLogin":       user.UserName != "",
		"user":          user,
		"cart":          cart,
		"address":       address,
		"totalPrice":    totalPrice,
		"order":         order,
		"orderItemList": orderItemList,
	})
}
