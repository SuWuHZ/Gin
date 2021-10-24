package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这一部分可以替换成从session/cookie中获取，
		username := c.Query("username")
		password := c.Query("password")

		if username == "abc" && password == "123" {
			c.JSON(http.StatusOK, gin.H{"message": "身份验证成功"})
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return
		}
	}
}
