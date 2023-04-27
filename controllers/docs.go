package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeSwaggerUI (c *gin.Context) {
	c.HTML(http.StatusOK, "swagger-ui.html", gin.H{
		"url": "http://localhost:3000/static/docs/openapi.yml",
	})
}