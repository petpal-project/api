package middleware

import (
	"net/http"
	"strings"

	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func EnsureValidToken(v *validator.Validator) gin.HandlerFunc {
	return func (c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := v.ValidateToken(c.Request.Context(), token)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return;
		}
		validatedClaims := claims.(*validator.ValidatedClaims)
		jwtSubject := validatedClaims.RegisteredClaims.Subject
		
		c.Set("user", jwtSubject)
		c.Next()
	}
}
