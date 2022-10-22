package middleware

import (
	"context"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func TokenAuth(client *auth.Client) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		headerToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), headerToken)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user", token.UID)
		c.Next()
	}
}