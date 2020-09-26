package main

import (
	"log"

	"github.com/testcontainers/helloworld/internal/server"
)

func main() {
	// Use a channel to wait indefinitely once running
	finish := make(chan bool)

	server.StartServing()

	log.Println("Ready, listening on 8080 and 8081")

	// Wait
	<-finish
}
