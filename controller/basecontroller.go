package controller

import (
	"net/http"
	"practice_gin_store/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (baseController BaseController) Success(ctx *gin.Context, msg string, url string) {
	ctx.HTML(http.StatusOK, "admin/base/success.html", gin.H{
		"message": msg,
		"url":     url,
	})
}

func (baseController BaseController) Fail(ctx *gin.Context, msg string, url string) {
	ctx.HTML(http.StatusOK, "admin/base/fail.html", gin.H{
		"message": msg,
		"url":     url,
	})
}

func (baseController BaseController) ChangeStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Request.FormValue("id"))
	table_name := ctx.Request.FormValue("table_name")
	status_value := ctx.Request.FormValue("status_value")
	if status_value == "0" {
		status_value = "1"
	} else {
		status_value = "0"
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "非法请求",
		})
	}
	if util.Db.Exec("update "+table_name+" set status="+status_value+" where id=?", id).Error != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "修改状态失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "修改状态成功",
		})
	}
}

func (baseController BaseController) ClearRedis(ctx *gin.Context) {
	_, err := util.Redis.Do("flushall")
	if err != nil {
		baseController.Fail(ctx, err.Error(), "index")
	} else {
		baseController.Success(ctx, "清除redis缓存成功", "index")
	}
}
