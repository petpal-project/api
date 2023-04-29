package controllers

import "github.com/gin-gonic/gin"

type PingService struct {}

func (s *PingService) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

