package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte(os.Getenv("1951e45dfd95b7bc77b68dd10621bb06869c0e4d87ea1bf5af2cd0d09d3b6cbf"))

// GenerateToken
func GenerateToken(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "test_case",
		Subject:   fmt.Sprint(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
