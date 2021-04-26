package main

import (
	"log"

	// coleta "github.com/elissonalvesilva/eng-zap-challenge-golang/routine/coleta"
	"github.com/elissonalvesilva/eng-zap-challenge-golang/routine/parser"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}
}

func main() {
	// coleta.InitColeta()
	// coleta.Run()
	parser.Run()
}
