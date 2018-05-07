package main

import (
	"log"
	"net/http"

	sw "./heroesservices"
)

//go:generate swagger generate spec
func main() {
	log.Printf("Heroes service is starting on port 8080...")

	router := sw.NewRouter()

	log.Printf("Heroes service up and running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
