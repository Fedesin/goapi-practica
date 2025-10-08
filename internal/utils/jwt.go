package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func secret() []byte {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		return []byte(v)
	}
	return []byte("claveultrasecreta")
}

func GenerateToken(email string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret())
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return secret(), nil
	})
}
