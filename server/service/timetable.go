package service

import (
	"ccsu/global"
	"ccsu/model"
	"ccsu/model/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTimetable(t request.TimeTableStruct) (err error, timetable *model.Timetable) {
	var tt model.Timetable
	if t.Weekly == 0 {
		err = global.DB.Where(&model.Timetable{Yearly: t.Yearly, StudentId: t.StuId}).Preload("Courses").First(&tt).Error
	} else {
		err = global.DB.Where(&model.Timetable{Yearly: t.Yearly, StudentId: t.StuId}).Preload("Courses", "find_in_set(?,weeks)", t.Weekly).Find(&tt).Error
	}
	return err, &tt
}

// 获取远程反向代理接口提供的课表
func GetRemoteTimeList(t request.TimeTableStruct) (error, *model.Timetable) {
	res, err := http.Get(fmt.Sprintf("%s?stuNum=%s&password=%s&yearly=%s&weekly=%s",
		global.Config.Remote.TimetableApi,
		t.StuNum,
		t.Password,
		t.Yearly,
		t.Weekly,
	))
	if err != nil {
		return err, nil
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			global.Log.Error(err.Error())
		}
	}()

	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	if err != nil {
		return err, nil
	}

	tt := &model.Timetable{}
	tt.Courses = make([]model.Course, 0)
	err = json.Unmarshal(data, &tt)
	if err != nil {
		return err, nil
	}
	return nil, tt
}

func SaveTimetable(t model.Timetable) (err error, tInter *model.Timetable) {
	err = global.DB.Create(&t).Error
	return err, &t
}