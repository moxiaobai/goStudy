package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/moxiaobai/goStudy/api/controllers"
	"net/http"
	"github.com/moxiaobai/goStudy/api/middleware"
)

func InitRoutes() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(middleware.Cors())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	//API Object routes
	apis := router.Group("/apis")
	{
		apis.GET("/", controllers.ListApiHandler)
		apis.GET("/:id", controllers.RetrieveApiHandler)
		apis.POST("/", controllers.AddApisHandler)
		apis.PATCH("/:id", controllers.UpdateApiHandler)
		apis.DELETE("/:id", controllers.DeleteApiHandler)
	}

	return router
}
