package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

//go:generate swagger generate spec
func main() {
	log.Printf("Heroes service is starting on port 8080...")

	router := NewRouter()

	log.Printf("Heroes service up and running on http://localhost:8080")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
