package v1

import (
	"ccsu/global"
	"ccsu/global/response"
	"ccsu/middleware"
	"ccsu/model"
	"ccsu/model/request"
	response2 "ccsu/model/response"
	"ccsu/service"
	"ccsu/utils/cst"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {

	var L request.LoginStruct

	_ = c.ShouldBindJSON(&L)
	err := global.Validator.Struct(L)
	if err != nil {
		response.FailWithValidatorError(err, c)
		return
	}

	if !store.Verify(L.CaptchaId, L.Captcha, true) {
		response.FailWithMessage(cst.ERROR_VERIFYCODE, c)
	} else {
		S := &model.Student{StuNum: L.StuNum, Password: L.Password}

		if err, stu := service.Login(S); err != nil {
			// 判断是否存在
			if e, isExist := service.IsExist(S.StuNum); e == nil && isExist {
				response.FailWithMessage(cst.ERROR_STU_OR_PWD, c)
			} else {
				err, ss := service.Register(*S)
				if err != nil {
					response.FailWithMessage(cst.REGISTER_FAIL, c)
				} else {
					tokenNext(c, ss)
				}
			}
		} else {
			tokenNext(c, *stu)
		}
	}
}

func GetStuInfo(c *gin.Context) {
	token := c.GetHeader("x-token")
	claim, _ := middleware.NewJWT().ParseToken(token)
	response.OkWithData(response2.StuInfoStruct{
		ID:     claim.ID,
		StuNum: claim.StuNum,
	}, c)
}

func tokenNext(c *gin.Context, stu model.Student) {
	j := &middleware.JWT{SigningKey: []byte(global.Config.Jwt.SigningKey)}

	clams := request.CustomClaims{
		ID:     stu.ID,
		StuNum: stu.StuNum,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7,
			Issuer:    "ygl",
		},
	}

	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage(cst.ERROR_GET_TOKEN_FAIL, c)
		return
	}

	response.OkWithData(response2.LoginResponse{
		Stu:       stu,
		Token:     token,
		ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
	}, c)
}
