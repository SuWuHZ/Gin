package controller

import (
	"log"
	"net/http"
	"time"

	"Gin/service"

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

func (BC *BaseController) ResponseFailed(c *gin.Context, code int) {

	message, ok := service.ErrorCodeMsgMap[code]
	if !ok {
		message = "failed."
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    nil,
		"message": message,
	})
}

func GetUserId(c *gin.Context) string {
	return c.Value("work_id").(string)
}

func (BC *BaseController) HandleHeartbeat(c *gin.Context) {
	log.Println("[HandleHeartbeat] timestamp:", time.Now().Unix())
	c.String(http.StatusOK, "connect ok")
}
