package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//添加API
func AddApisHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"errcode" : 0,
		"errmsg"  : "ok",
	})
}

//更新API
func UpdateApiHandler(c *gin.Context) {

}

//删除API
func DeleteApiHandler(c *gin.Context) {

}

//列出API
func ListApiHandler(c *gin.Context) {

}

//检索API
func RetrieveApiHandler(c *gin.Context) {

}
