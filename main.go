package main

import (
	"pet-pal/api/config"
	"pet-pal/api/middleware"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	authClient := config.InitFirebase()

	r := gin.Default()
	r.GET("/ping", pingHandler)

	api := r.Group("/api")
	{
		api.Use(middleware.TokenAuth(authClient))
		users := api.Group("/users")
		{
			users.GET("/")
			users.POST("/")
		}
		cats := api.Group("/cats")
		{
			breeds := cats.Group("/breeds")
			{
				breeds.GET("/")
			}
		}
	}

	r.Run(":3000")
}
