package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaindy7633/package-gin-framework/bootstrap"
	"github.com/kaindy7633/package-gin-framework/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	_ = r.Run(":" + global.App.Config.App.Port)
}
