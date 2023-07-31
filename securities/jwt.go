package securities

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// GeneratesJWT generates JSON Web Token for authentication purpose
func GenerateJWT(id uuid.UUID, username string) (string, error) {
	JWTKeySecret := []byte(os.Getenv("JWT_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(12 * time.Hour)
	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = username

	tokenString, err := token.SignedString(JWTKeySecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
