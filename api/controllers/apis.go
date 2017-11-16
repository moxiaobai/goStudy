package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "github.com/moxiaobai/goStudy/api/models"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	//"go.uber.org/zap"
	"log"
)

//添加API
func AddApisHandler(c *gin.Context) {
	var api    Api
	var err    error
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		err = c.ShouldBindJSON(&api)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&api, binding.Form)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := api.Add()
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id" : id,})
}

//更新API
func UpdateApiHandler(c *gin.Context) {
	cid := c.Param("id")
	//转换为整型
	id, _ := strconv.Atoi(cid)

	api := Api{Id: id}
	if err := c.ShouldBindJSON(&api); err == nil {
		numbers, _ := api.Update()

		c.JSON(http.StatusOK, gin.H{"numbers" : numbers,})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//删除API
func DeleteApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	api := Api{Id: id}
	numbers, _ :=  api.Delete()

	c.JSON(http.StatusOK, gin.H{"numbers" : numbers,})
}

//列出API
func ListApiHandler(c *gin.Context) {
	cSize   := c.DefaultQuery("size", "2")
	cOffset := c.DefaultQuery("offset", "1")


	//logger, _ := zap.NewProduction()
	//defer logger.Sync()
	//logger.Info("failed to fetch URL", )

	size, _  := strconv.Atoi(cSize)
	offset,_ := strconv.Atoi(cOffset)

	var api Api
	apis,_ := api.List(offset, size)

	c.JSON(http.StatusOK, gin.H{"apis" : apis,})
}

//检索API
func RetrieveApiHandler(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	log.Print("hello")

	a := Api{Id: id}
	api, err  := a.Retrieve()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"api" : nil,})
		return
	}

	c.JSON(http.StatusOK, gin.H{"api" : api,})
}
