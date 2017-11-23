package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SetTokenHandler(c *gin.Context) {
	expireToken  := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)

	claims := MyCustomClaims{
		"moxiaobai",
		jwt.StandardClaims{
			ExpiresAt:expireToken,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires:expireCookie, HttpOnly:true}
	http.SetCookie(c.Writer, &cookie)
}
