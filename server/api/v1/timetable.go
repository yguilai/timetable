package v1

import (
	"ccsu/global/response"
	"ccsu/model/request"
	"ccsu/service"
	"ccsu/utils/cst"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func TimeTable(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sid"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, stu := service.GetStudentById(uint(id))
	if err != nil || stu == nil {
		response.FailWithMessage(cst.ERROR_STU_NOT_EXIST, c)
		return
	}

	var t request.TimeTableStruct
	_ = c.ShouldBindQuery(&t)

	t.StuId = stu.ID
	err, timetable := service.GetTimetable(t)
	if err != nil && err != gorm.ErrRecordNotFound {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if timetable.ID != 0 {
		response.OkWithData(timetable, c)
	} else {
		t.StuNum = stu.StuNum
		t.Password = stu.Password
		err, tt := service.GetRemoteTimeList(t)
		if err != nil {
			response.FailWithMessage(cst.ERROR_PROXY_TT_FAIL, c)
			return
		}
		tt.StudentId = stu.ID
		err, _ = service.SaveTimetable(*tt)
		if err != nil {
			response.FailWithMessage(cst.ERROR_SAVE_TT_FAIL, c)
			return
		}

		err, r := service.GetTimetable(t)
		if err != nil {
			response.FailWithMessage(cst.ERROR_QUERY_TT_FAIL, c)
			return
		}
		response.OkWithData(r, c)
	}
}
