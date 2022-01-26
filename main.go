package main

import (
	"log"
	"github.com/danilomarques1/gopmserver/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	s.Init()
	s.Start()
}
