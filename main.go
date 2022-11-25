package main

import (
	"log"

	"github.com/Fachrulmustofa20/go-microservice-user/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Init()

	err = cfg.Start()
	if err != nil {
		log.Fatal(err)
	}
}
