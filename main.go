package main

import (
	"log"
	"net/http"

	"github.com/cmumford/go-starter.git/api"
)

func main() {
	http.HandleFunc("/", api.RootHandler)
	http.HandleFunc("/health", api.HealthHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
