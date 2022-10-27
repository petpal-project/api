package main

import (
	"pet-pal/api/config"
	"pet-pal/api/controllers"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	authClient := config.InitFirebase()
	var DB *gorm.DB = config.InitDb()

	var r *gin.Engine = gin.Default()
	r.GET("/ping", pingHandler)

	var api *gin.RouterGroup = r.Group("/api")
	{
		api.Use(middleware.TokenAuth(authClient, DB))
		var users *gin.RouterGroup = api.Group("/users")
		{
			users.GET("/", controllers.GetUser)
			users.POST("/", controllers.PostUser)
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
