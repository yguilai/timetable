package request

type LoginStruct struct {
	StuNum    string `json:"stuNum" validate:"required,len=12" comment:"学号"`
	Password  string `json:"password" validate:"required,len=6" comment:"密码"`
	CaptchaId string `json:"captchaId" validate:"required"`
	Captcha   string `json:"captcha" validate:"required" comment:"验证码"`
}
