package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeSwaggerUI (c *gin.Context) {
	c.HTML(http.StatusOK, "swagger-ui.html", gin.H{
		"url": fmt.Sprintf("http://%s/static/docs/openapi.yml", c.Request.Host),
	})
}