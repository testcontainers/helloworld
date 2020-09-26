package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/testcontainers/helloworld/internal/util"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG")
}

func uuidHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, uuid.New().String())
}

func StartServing() {
	fs := http.FileServer(http.Dir("./static"))

	server8080 := http.NewServeMux()
	server8080.Handle("/", fs)
	server8080.HandleFunc("/ping", pingHandler)
	server8080.HandleFunc("/uuid", uuidHandler)

	server8081 := http.NewServeMux()
	server8081.Handle("/", fs)
	server8081.HandleFunc("/ping", pingHandler)
	server8081.HandleFunc("/uuid", uuidHandler)

	// Delay, if configured to do so with an environment variable
	delayStart := util.GetEnvInt("DELAY_START_MSEC", 0)
	log.Printf("DELAY_START_MSEC: %d\n", delayStart)

	// Delay once before the server on port 8080 starts
	log.Printf("Sleeping for %d ms", delayStart)
	time.Sleep(time.Duration(delayStart) * time.Millisecond)

	log.Println("Starting server on port 8080")
	go func() {
		log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, server8080)))
	}()

	// Delay again before the server on port 8081 starts
	log.Printf("Sleeping for %d ms", delayStart)
	time.Sleep(time.Duration(delayStart) * time.Millisecond)

	log.Println("Starting server on port 8081")
	go func() {
		log.Fatal(http.ListenAndServe(":8081", handlers.LoggingHandler(os.Stdout, server8081)))
	}()
}
