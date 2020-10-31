package router

import (
	v1 "ccsu/api/v1"
	"ccsu/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuth(group *gin.RouterGroup) {
	auth := group.Group("/auth")
	{
		auth.GET("/captcha", v1.Captcha)
		auth.POST("/login", v1.Login)
		auth.GET("/info", v1.GetStuInfo).Use(middleware.JwtAuth())
	}
}
