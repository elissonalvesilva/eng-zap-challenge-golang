package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	server "github.com/elissonalvesilva/eng-zap-challenge-golang/api/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}
}

func runAPI() {
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

func checkLockFile(attemptsLockFile *int, stopCheck *bool) {
	lockfile := os.Getenv("PATH_DADOS") + "lock"
	catalog := os.Getenv("PATH_DADOS") + os.Getenv("FILENAME_PARSED_CATALOG")

	if _, err := os.Stat(catalog); err == nil {
		if _, err := os.Stat(lockfile); err != nil {
			*stopCheck = true
		}
	}
	*attemptsLockFile++
}

func main() {
	attemptsLockFile := 0
	stopCheck := false
	for t := range time.Tick(2 * time.Second) {
		checkLockFile(&attemptsLockFile, &stopCheck)
		fmt.Println("Check lock file at: ", t.Local(), "Attempts: ", attemptsLockFile)
		if stopCheck {
			break
		}
	}

	fmt.Println("Lock file not exists, now can run API ")
	runAPI()
}
