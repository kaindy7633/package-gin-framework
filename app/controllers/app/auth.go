package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaindy7633/package-gin-framework/app/common/request"
	"github.com/kaindy7633/package-gin-framework/app/common/response"
	"github.com/kaindy7633/package-gin-framework/app/services"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.GuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
