package controller

import (
	"Gin/models"
	"Gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	*BaseController
}

func (C *TestController) HandleGetUserInfo(c *gin.Context) {
	workId := c.Query("work_id")

	errCode, data := service.GetUserInfo(workId)
	if errCode == 0 {
		C.ResponseSuccess(c, data)
		return
	} else {
		C.ResponseFailed(c, errCode)
	}
}

func (C *TestController)HandleAddUserInfo(c *gin.Context) {
	workId := GetUserId(c)

	var postJson models.User

	if err := c.ShouldBindJSON(&postJson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"message": "request body error.",
		})
		return
	}
	if workId != postJson.WordID {
		c.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"message": "request body error.",
		})
		return
	}

	if errCode, data := service.AddOneUser(&postJson); errCode != 0 {
		C.ResponseFailed(c, errCode)
	}else {
		C.ResponseSuccess(c, data)
	}
	return
}
