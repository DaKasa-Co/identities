package securities

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Authenticate - Middleware to authenticate requests with
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		CoreAuthenticate(c)
	}
}

// CORSMiddleware -
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://"+os.Getenv("MFE_DOMAIN")+":"+os.Getenv("MFE_PORT"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CoreAuthenticate(c *gin.Context) {
	token, err := jwt.Parse(c.GetHeader("X-JWT"), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized", "detail": err.Error()})
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized", "detail": jwt.ErrInvalidKey.Error()})
}
