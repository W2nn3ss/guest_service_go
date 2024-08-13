package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func CheckBearerToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != os.Getenv("API_TOKEN") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
