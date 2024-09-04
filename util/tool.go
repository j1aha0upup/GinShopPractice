package util

import (
	"html/template"
	"time"

	"github.com/russross/blackfriday/v2"
)

func Unix2Time(unix int) string {
	return time.Unix(int64(unix), 0).String()
}

func Markdown2Html(temp string) template.HTML {
	return template.HTML(string(blackfriday.Run([]byte(temp))))
}

func GoodsPriceMulGoodsNum(price float64, num int) float64 {
	return price * float64(num)
}
func TitleStrCut(temp string, cutNum int) string {
	if len(temp) > cutNum {
		temp = temp[0:cutNum]
	}
	return temp
}
