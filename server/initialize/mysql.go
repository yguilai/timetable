package initialize

import (
	"ccsu/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	c := global.Config.Mysql

	if db, err := gorm.Open(mysql.Open(c.Username+":"+c.Password+"@("+c.Path+")/"+c.Dbname+"?"+c.Config), &gorm.Config{}); err != nil {
		global.Log.Error(err.Error())
	} else {
		d, _ := db.DB()
		d.SetMaxOpenConns(c.MaxOpenConns)
		d.SetMaxIdleConns(c.MaxIdleConns)
		global.DB = db
	}
}
