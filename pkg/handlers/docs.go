package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocHandler struct {}

func (h DocHandler) ServeSwaggerUI(c *gin.Context) {
	c.HTML(http.StatusOK, "swagger-ui.html", gin.H{
		"url": fmt.Sprintf("http://%s/static/docs/openapi.yml", c.Request.Host),
	})
}