package main

import (
	"github.com/kaindy7633/package-gin-framework/bootstrap"
	"github.com/kaindy7633/package-gin-framework/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化校验器
	bootstrap.InitializeValidator()

	// 启动服务器
	bootstrap.RunServer()
}
