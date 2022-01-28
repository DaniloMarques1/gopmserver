package main

import (
	"log"

	"github.com/danilomarques1/gopmserver/server"

	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	Id string
	jwt.Claims
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	s.Init()
	s.Start()
}
