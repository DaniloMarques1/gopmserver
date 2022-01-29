package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Id string
	jwt.StandardClaims
}

const TOKEN_EXPIRATION_TIME = 3600

// receive the users id and return a token and a possible error
func GenToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + TOKEN_EXPIRATION_TIME,
			Issuer:    "gopmserver",
		},
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// returns an error if token is not valid
func VerifyToken(tokenStr string) (string, error) {
	var tokenClaims TokenClaims
	_, err := jwt.ParseWithClaims(tokenStr, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("ERR signing method\n")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	return tokenClaims.Id, nil
}

// will receive the bearer token like Bearer {token}
// and return only the token part
func GetTokenFromHeader(authHeader string) (string, error) {
	strSlice := strings.Split(authHeader, " ")
	if len(strSlice) < 2 {
		return "", fmt.Errorf("No token provided")
	}
	return strSlice[1], nil
}
