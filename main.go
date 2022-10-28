package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	// authClient := config.InitFirebase()
	config.InitDb()

	var r *gin.Engine = gin.Default()
	r.GET("/ping", pingHandler)

	var api *gin.RouterGroup = r.Group("/api")
	{
		// api.Use(middleware.TokenAuth(authClient, DB))
		api.Use(middleware.TempUserAuth)
		var users *gin.RouterGroup = api.Group("/users")
		{
			users.GET("/", controllers.GetUser)
			users.POST("/", controllers.PostUser)
			users.DELETE("/", controllers.DeleteUser)
		}
		var cats *gin.RouterGroup = api.Group("/cats")
		{
			var breeds *gin.RouterGroup = cats.Group("/breeds")
			{
				breeds.GET("/")
			}
		}
	}

	r.Run(":3000")
}
