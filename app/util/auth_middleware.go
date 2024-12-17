// app/util/auth_middleware.go
package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 是 JWT 身份验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the "Authorization" header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix from the token string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Validate the token
		username, err := ValidateToken(tokenString)
		if err != nil || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set the user information in the context (optional)
		c.Set("username", username)

		c.Next()
	}
}
