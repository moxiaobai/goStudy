package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/moxiaobai/goStudy/api/models"
	"strconv"
	"log"
)

//添加API
func AddApisHandler(c *gin.Context) {
	var api models.Api
	if c.Bind(&api) != nil {
		c.JSON(404, gin.H{"JSON=== status": "binding JSON error!"})
		return
	}

	db   := models.InitDb()
	id := db.AddApis(api)

	c.JSON(http.StatusOK, gin.H{
		"id" : id,
	})
}

//更新API
func UpdateApiHandler(c *gin.Context) {
	cid := c.Param("id")
	//转换为整型
	id, _ := strconv.Atoi(cid)

	var api models.Api
	err := c.BindJSON(&api)
	if err != nil {
		log.Println(err)
	}

	db   := models.InitDb()
	numbers := db.UpdateApi(id, api)

	c.JSON(http.StatusOK, gin.H{
		"numbers" : numbers,
	})
}

//删除API
func DeleteApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	db   := models.InitDb()
	numbers :=  db.DeleteApi(id)

	c.JSON(http.StatusOK, gin.H{
		"numbers" : numbers,
	})
}

//列出API
func ListApiHandler(c *gin.Context) {
	db   := models.InitDb()
	apis := db.ListApi()

	c.JSON(http.StatusOK, gin.H{
		"apis" : *apis,
	})
}

//检索API
func RetrieveApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	db   := models.InitDb()
	api  := db.RetrieveApi(id)

	c.JSON(http.StatusOK, gin.H{
		"apis" : api,
	})
}