package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return
		} else {
			c.Set("work_id", token)
			c.Next()
		}
	}
}
