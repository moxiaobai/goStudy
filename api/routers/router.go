package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/moxiaobai/goStudy/api/controllers"
	"net/http"
)

func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	//API Object routes
	apis := router.Group("/apis")
	{
		apis.GET("/", controllers.ListApiHandler)
		apis.GET("/:name", controllers.RetrieveApiHandler)
		apis.POST("/", controllers.AddApisHandler)
		apis.PATCH("/:name", controllers.UpdateApiHandler)
		apis.DELETE("/:name", controllers.DeleteApiHandler)
	}

	return router
}
