package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("JWT_SECRET")

func GenerateJWT(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}
