package main

import (
	"ccsu/core"
	"ccsu/global"
	"ccsu/initialize"
)

func main() {
	// 初始化数据库连接
	initialize.Mysql()
	d, _ := global.DB.DB()
	initialize.DBTables()
	// 释放数据库资源
	defer d.Close()
	core.RunServer()
}
