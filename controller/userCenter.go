package controller

import (
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserCenterController struct {
	BaseController
}

func (userCenterController UserCenterController) UserCenter(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	temp := []model.Order{}
	var waitPay, waitAccept int64
	util.Db.Where("uid=? and order_status=?", user.Id, 0).Find(&temp).Count(&waitPay)
	util.Db.Where("uid=? and order_status=?", user.Id, 3).Find(&temp).Count(&waitAccept)

	ctx.HTML(http.StatusOK, "shop/user/user.html", gin.H{
		"topNav":     topNav,
		"topNavLen":  len(topNav) - 1,
		"middleNav":  middleNav,
		"footerNav":  footerNav,
		"goodsCate":  goodsCate,
		"focus":      focus,
		"isLogin":    user.UserName != "",
		"user":       user,
		"pathname":   ctx.Request.URL.Path,
		"waitAccept": waitAccept,
		"waitPay":    waitPay,
	})
}

func (userCenterController UserCenterController) UserOrder(ctx *gin.Context) {
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	orderList := []model.Order{}
	orderItemList := []model.OrderItem{}

	orderStatus, err := strconv.Atoi(ctx.Request.FormValue("orderStatus"))
	if err != nil {
		orderStatus = -1
	}
	currentPage, _ := strconv.Atoi(ctx.Request.FormValue("currentPage"))
	page := model.Pagination{
		PageSize:     5,
		CurrentPage:  currentPage,
		VisiblePages: 8,
	}

	if page.CurrentPage != 0 {
		page.CurrentPage--
	}

	searchValue := strings.TrimSpace(ctx.Request.FormValue("searchValue"))
	whereStr := ""
	OrderIdList := []int{}
	if searchValue != "" {
		whereStr += "uid=" + strconv.Itoa(user.Id) + " and goods_title like \"%" + searchValue + "%\""
		util.Db.Where(whereStr).Find(&orderItemList)
		for i := 0; i < len(orderItemList); i++ {
			OrderIdList = append(OrderIdList, orderItemList[i].OrderId)
		}
		if orderStatus < 0 {
			util.Db.Where("id in ?", OrderIdList).Find(&orderList)
		} else {
			util.Db.Where("order_status=? and id in ?", orderStatus, OrderIdList).Find(&orderList)
		}
	} else {
		if orderStatus < 0 {
			util.Db.Where("uid=?", user.Id).Find(&orderList)
		} else {
			util.Db.Where("uid=? and order_status=?", user.Id, orderStatus).Find(&orderList)
		}
	}

	if len(orderList)%page.PageSize == 0 {
		page.TotalPages = len(orderList) / page.PageSize
	} else {
		page.TotalPages = len(orderList)/page.PageSize + 1
	}

	if searchValue != "" {
		if orderStatus < 0 {
			util.Db.Where("id in ?", OrderIdList).Limit(page.PageSize).Offset(page.CurrentPage * page.PageSize).Preload("OrderItem").Find(&orderList)
		} else {
			util.Db.Where("order_status=? and id in ?", orderStatus, OrderIdList).Limit(page.PageSize).Offset(page.CurrentPage * page.PageSize).Preload("OrderItem").Find(&orderList)
		}
	} else {
		if orderStatus < 0 {
			util.Db.Where("uid=?", user.Id).Limit(page.PageSize).Offset(page.CurrentPage * page.PageSize).Preload("OrderItem").Find(&orderList)
		} else {
			util.Db.Where("uid=? and order_status=?", user.Id, orderStatus).Limit(page.PageSize).Offset(page.CurrentPage * page.PageSize).Preload("OrderItem").Find(&orderList)
		}
	}

	page.CurrentPage++

	ctx.HTML(http.StatusOK, "shop/user/order.html", gin.H{
		"topNav":      topNav,
		"topNavLen":   len(topNav) - 1,
		"middleNav":   middleNav,
		"footerNav":   footerNav,
		"goodsCate":   goodsCate,
		"focus":       focus,
		"isLogin":     user.UserName != "",
		"user":        user,
		"pathname":    ctx.Request.URL.Path,
		"orderStatus": orderStatus,
		"page":        page,
		"orderList":   orderList,
		"order":       len(orderList) != 0,
		"searchValue": searchValue,
	})
}

func (userCenterController UserCenterController) UserOrderInfo(ctx *gin.Context) {
	order := model.Order{}
	topNav, middleNav, footerNav, goodsCate, focus, user := getPublicInfo(ctx)
	order.Id, _ = strconv.Atoi(ctx.Request.FormValue("orderId"))
	util.Db.Where("id=?", order.Id).Preload("OrderItem").Find(&order)

	ctx.HTML(http.StatusOK, "shop/user/orderInfo.html", gin.H{
		"topNav":    topNav,
		"topNavLen": len(topNav) - 1,
		"middleNav": middleNav,
		"footerNav": footerNav,
		"goodsCate": goodsCate,
		"focus":     focus,
		"isLogin":   user.UserName != "",
		"user":      user,
		"pathname":  ctx.Request.URL.Path,
		"order":     order,
	})
}
