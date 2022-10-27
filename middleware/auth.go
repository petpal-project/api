package middleware

import (
	"context"
	"pet-pal/api/controllers"
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
		var userId uint = controllers.GetUserIdFromFirebaseId(token.UID, db)
		if userId == 0 {
			c.JSON(500, gin.H{"error": "User associated with the token does not exist in database"})
		}
		c.Set("user", userId)
		c.Next()
	}
}
