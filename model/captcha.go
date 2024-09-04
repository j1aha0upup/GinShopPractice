package model

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type MyCaptcha struct {
	Id   string `json:"captchaId"`
	B64s string `json:"captchaB64s"`
}

var store = base64Captcha.DefaultMemStore

// base64Captcha.Driver is interface but base64Captcha.DriverString is struct
var driver base64Captcha.Driver

func (mycaptcha *MyCaptcha) CreateCaptcha() {
	driverString := base64Captcha.DriverString{Height: 40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3, G: 102, B: 214, A: 125},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()

	if err != nil {
		panic("验证码生成失败！")
	} else {
		mycaptcha.Id = id
		mycaptcha.B64s = b64s
	}
}

func VerifyCaptcha(id string, answer string) bool {
	return store.Verify(id, answer, true)
}
