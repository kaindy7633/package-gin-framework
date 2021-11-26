package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaindy7633/package-gin-framework/app/controllers/app"
	"github.com/kaindy7633/package-gin-framework/app/middleware"
	"github.com/kaindy7633/package-gin-framework/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.GuardName))
	{
		authRouter.GET("/auth/info", app.Info)
	}
}
