package main

import (
	"log"

	server "github.com/elissonalvesilva/eng-zap-challenge-golang/api/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}
}

func main() {
	server.NewApp().Run()
}
