package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaindy7633/package-gin-framework/app/common/request"
	"github.com/kaindy7633/package-gin-framework/app/common/response"
	"github.com/kaindy7633/package-gin-framework/app/services"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
