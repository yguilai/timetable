package v1

import (
	"ccsu/global"
	"ccsu/global/response"
	res "ccsu/model/response"
	"ccsu/utils/cst"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context)  {
	cg := global.Config.Captcha
	driver := base64Captcha.NewDriverDigit(cg.Height, cg.Width, cg.Length, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage(cst.VERIFY_GENERATE_FAIL , c)
	} else {
		response.OkDetailed(res.CaptchaResponse{
			CaptchaId: id,
			PicPath: b64s,
		}, "验证码获取成功", c)
	}
}
