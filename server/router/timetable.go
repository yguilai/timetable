package router

import (
	v1 "ccsu/api/v1"
	"ccsu/middleware"
	"github.com/gin-gonic/gin"
)

func InitTimetable(Router *gin.RouterGroup)  {
	Timetable := Router.Group("/stu/:sid/timetable").Use(middleware.JwtAuth())
	{
		Timetable.GET("", v1.TimeTable)
	}
}
