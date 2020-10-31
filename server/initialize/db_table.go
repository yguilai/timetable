package initialize

import (
	"ccsu/global"
	"ccsu/model"
)

func DBTables() {
	db := global.DB

	db.AutoMigrate(
		model.Student{},
		model.Course{},
		model.Timetable{},
	)
	global.Log.Debug("register table success")
}
