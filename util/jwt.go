package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Id string
	jwt.Claims
}

const TOKEN_EXPIRATION = 3600

func GenToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		Id: id,
		Claims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + TOKEN_EXPIRATION,
			Issuer:    "gopmserver",
		},
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
