package main

import (
	"log"

	"github.com/danilomarques1/gopmserver/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading env variable %v\n", err)
	}

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	s.Init()
	s.Start()
}
