package controller

import (
	"Gin/service"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	*BaseController
}

func (TC *TestController) HandleGetUserInfo(c *gin.Context) {
	userId := GetUserId(c)

	errCode, data := service.GetUserInfo(userId)
	if errCode == 0 {
		TC.ResponseSuccess(c, data)
		return
	} else {
		TC.ResponseFailed(c, errCode)
	}
}
