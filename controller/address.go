package controller

import (
	"net/http"
	"practice_gin_store/model"
	"practice_gin_store/util"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	BaseController
}

func (addressController AddressController) AddAddress(ctx *gin.Context) {
	address, addressList := model.Address{}, []model.Address{}
	user := model.User{}
	userTemp, ok := ctx.Get("user")
	if ok {
		user = userTemp.(model.User)
	}
	util.Db.Where("uid=?", user.Id).Find(&addressList)

	if len(addressList) > 10 {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "收货地址超过十个",
		})
	} else {
		address.Name = ctx.PostForm("name")
		address.Phone = ctx.PostForm("phone")
		address.Address = ctx.PostForm("address")
		address.AddTime = int(time.Now().Unix())
		address.Uid = user.Id
		address.DefaultAddress = 1
		for i := 0; i < len(addressList); i++ {
			addressList[i].DefaultAddress = 0
		}
		util.Db.Save(&addressList)
		util.Db.Create(&address)
		util.Db.Where("uid=?", user.Id).Order("id desc").Find(&addressList)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  addressList,
		})
	}
}

func (addressController AddressController) EditAddress(ctx *gin.Context) {
	address, addressList := model.Address{}, []model.Address{}
	user := model.User{}
	userTemp, ok := ctx.Get("user")
	if ok {
		user = userTemp.(model.User)
	}
	util.Db.Where("uid=?", user.Id).Find(&addressList).Updates(map[string]interface{}{"default_address": 0})
	address.Id, _ = strconv.Atoi(ctx.PostForm("id"))
	util.Db.Where("id=?", address.Id).Find(&address)
	address.Name = ctx.PostForm("name")
	address.Phone = ctx.PostForm("phone")
	address.Address = ctx.PostForm("address")
	address.DefaultAddress = 1
	util.Db.Save(&address)

	util.Db.Where("uid=?", user.Id).Order("id desc").Find(&addressList)
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  addressList,
	})
}

func (addressController AddressController) GetOneAddressList(ctx *gin.Context) {
	address := model.Address{}
	addressId := ctx.Request.FormValue("addressId")
	util.Db.Where("id=?", addressId).Find(&address)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  address,
	})
}

func (addressController AddressController) ChangeDefaultAddress(ctx *gin.Context) {
	addressId := ctx.Request.FormValue("addressId")
	addressList, address := []model.Address{}, model.Address{}
	user := model.User{}
	userTemp, ok := ctx.Get("user")
	if ok {
		user = userTemp.(model.User)
	}
	util.Db.Where("uid=?", user.Id).Find(&addressList).Updates(map[string]interface{}{"default_address": 0})
	util.Db.Where("id=?", addressId).Find(&address).Updates(map[string]interface{}{"default_address": 1})

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
