package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func TempUserAuth(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	userId, err := strconv.Atoi(authHeader)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Auth header must be a userId (int)",
		})
		c.Abort()
		return
	}
	c.Set("user", userId)
	c.Next()
}
