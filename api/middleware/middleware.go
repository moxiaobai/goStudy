/*
中间件
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"github.com/moxiaobai/goStudy/api/controllers"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("Auth")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
			})

			c.Abort()
			return
		}

		splitCookie := strings.Split(cookie.String(), "Auth=")

		token, err := jwt.ParseWithClaims(splitCookie[1], &controllers.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			}
		})

		if claims, ok := token.Claims.(*controllers.MyCustomClaims); ok && token.Valid {
			c.Set("claims", claims)
		}
	}
}
