package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingService struct{}

func (service *PingService) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
