package routes

import (
	"github.com/gin-gonic/gin"
	"golang-gin/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		user := main.Group("user")
		{
			user.GET("/", controllers.Hello)
			user.GET("/all", controllers.GetAllUsers) 
		}
	}

	return router
}
