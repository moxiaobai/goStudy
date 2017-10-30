/*
中间件
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error" : "Unauthorized",
		})

		c.Abort()
		return
	}
}
