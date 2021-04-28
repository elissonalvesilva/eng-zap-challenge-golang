package main

import (
	"log"
	"sync"

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
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		server.NewApp(4513).Run("Platform API")
		wg.Done()
	}()

	go func() {
		server.NewApp(4514).Run("Platform API")
		wg.Done()
	}()

	wg.Wait()
}
