package middleware

import (
	"elotus/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// TokenAuthentication validates the token and injects user info into context
func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"meta": gin.H{
					"code":    401000,
					"message": "Unauthorized: Missing or invalid token",
				},
			})
			return
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := jwt.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"meta": gin.H{
					"code":    401000,
					"message": "Unauthorized: Invalid token",
				},
			})
			return
		}

		// Inject user information into context
		c.Set("username", claims.Username)
		c.Next()
	}
}
