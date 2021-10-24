package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (BC *BaseController) ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    data,
		"message": "success",
	})
}

func (BC *BaseController) HandleHeartbeat(c *gin.Context) {
	log.Println("[HandleHeartbeat] timestamp:", time.Now().Unix())
	c.String(http.StatusOK, "ok")
}
