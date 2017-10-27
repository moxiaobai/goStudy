package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "github.com/moxiaobai/goStudy/api/models"
	"strconv"
)

//添加API
func AddApisHandler(c *gin.Context) {
	var a Api
	if err := c.ShouldBindJSON(&a); err == nil {
		id, _ := a.Add()

		c.JSON(http.StatusOK, gin.H{"id" : id,})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//更新API
func UpdateApiHandler(c *gin.Context) {
	cid := c.Param("id")
	//转换为整型
	id, _ := strconv.Atoi(cid)

	a := Api{Id: id}
	if err := c.ShouldBindJSON(&a); err == nil {
		numbers, _ := a.Update()

		c.JSON(http.StatusOK, gin.H{"numbers" : numbers,})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//删除API
func DeleteApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	a := Api{Id: id}
	numbers, _ :=  a.Delete()

	c.JSON(http.StatusOK, gin.H{"numbers" : numbers,})
}

//列出API
func ListApiHandler(c *gin.Context) {
	cSize   := c.DefaultQuery("size", "2")
	cOffset := c.DefaultQuery("offset", "1")

	size, _  := strconv.Atoi(cSize)
	offset,_ := strconv.Atoi(cOffset)

	var a Api
	apis,_ := a.List(offset, size)

	c.JSON(http.StatusOK, gin.H{"apis" : apis,})
}

//检索API
func RetrieveApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	a := Api{Id: id}
	api, err  := a.Retrieve()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"api" : nil,})
		return
	}

	c.JSON(http.StatusOK, gin.H{"api" : api,})
}
