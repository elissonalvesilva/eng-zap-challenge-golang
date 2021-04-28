package main

import (
	"log"

	"github.com/elissonalvesilva/eng-zap-challenge-golang/api"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/routine"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}
}

func main() {
	routine.Run()
	api.Run()
}
