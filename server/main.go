package main

import (
	"quickSearch/core"
	"quickSearch/global"
	"quickSearch/initialize"
)

func main() {
	//初始化配置信息
	global.CONFIG = core.Config()
	//初始化数据库
	global.DB = initialize.GormSqlite()
	initialize.RegisterTables() // 初始化表
	// 程序结束前关闭数据库链接
	db := global.DB
	defer db.Close()
	core.RunWindowsServer()
}
