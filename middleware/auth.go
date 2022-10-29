package middleware

import (
	"context"
	"pet-pal/api/controllers"
	"strconv"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TokenAuth(client *auth.Client, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		headerToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), headerToken)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		var userId uint = controllers.GetUserIdFromFirebaseId(token.UID)
		c.Set("user", userId)
		c.Next()
	}
}

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
