package main

import (
	server "github.com/elissonalvesilva/eng-zap-challenge-golang/api/server"
)

func main() {
	server.NewApp().Run()
}
