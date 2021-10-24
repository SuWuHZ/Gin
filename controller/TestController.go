package controller

import "github.com/gin-gonic/gin"

type TestController struct {
	*BaseController
}

func (TC *TestController) HandleGetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
