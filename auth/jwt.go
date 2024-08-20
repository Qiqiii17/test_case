package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecretKey = []byte("1951e45dfd95b7bc77b68dd10621bb06869c0e4d87ea1bf5af2cd0d09d3b6cbf")

// GenerateToken
func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecretKey)

	return token, err
}

type Claims struct {
	IdUser uint
	jwt.StandardClaims
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
