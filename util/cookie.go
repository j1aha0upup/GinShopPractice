package util

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetCookie(ctx *gin.Context, cookieName string, cookieValue string, maxAge int) {
	//!!!must len == 8
	tempcookieValue, err := DesEncrypt([]byte(cookieValue), []byte("12345678"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		cookieValue = string(tempcookieValue)
	}

	ctx.SetCookie(cookieName, cookieValue, maxAge, "/", ctx.Request.URL.Host, false, true)
}

func GetCookie(ctx *gin.Context, cookieName string, cookieValue interface{}) {
	temp, _ := ctx.Cookie(cookieName)

	if temp != "" {
		tempCookieValue, err := DesDecrypt([]byte(temp), []byte("12345678"))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			json.Unmarshal([]byte(tempCookieValue), cookieValue)
		}
	}
}
