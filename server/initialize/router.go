package initialize

import (
	"ccsu/global"
	"ccsu/middleware"
	"ccsu/router"
	"github.com/gin-gonic/gin"
	"time"
)

func Routers() *gin.Engine {
	var r *gin.Engine

	if global.Config.System.Mode == "dev" {
		r = gin.Default()
	} else {
		r = gin.New()
		// 使用zap接收gin的访问日志
		r.Use(
			middleware.LoggerWithZap(global.Log, time.RFC3339, true),
			middleware.RecoveryWithZap(global.Log, true),
		)
	}

	// 支持https
	//r.Use(middleware.LoadTls())
	// 跨域处理
	r.Use(middleware.Cors())

	// 统一添加路由组前缀
	ApiGroup := r.Group("")
	//router.InitTeacherEvaluation(ApiGroup)
	router.InitTimetable(ApiGroup)
	router.InitAuth(ApiGroup)
	return r
}
