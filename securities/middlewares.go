package securities

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Authenticate - Middleware to authenticate requests with
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		CoreAuthenticate(c)
	}
}

func CoreAuthenticate(c *gin.Context) {
	apiKey := c.GetHeader("X-API-Key")
	if apiKey == os.Getenv("API-KEY") {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized,
		gin.H{"error": "Unauthorized"})
}
