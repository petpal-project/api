package main

import (
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)

	api := r.Group("/api")
	{
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

	r.Run(":3000");
}