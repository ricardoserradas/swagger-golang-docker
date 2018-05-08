package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetHero returns a Person with a Hero name, given first and last name
func GetHero(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /api/Heroes/Get/{firstName}/{lastName} getHero
	//
	// Returns the Person with a Hero name assigned
	//
	// Could be any Hero
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: firstName
	//   in: path
	//   description: the first name of the Person
	//   required: true
	//   type: string
	// - name: lastName
	//   in: path
	//   description: the last name of the Person
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Hero response

	params := mux.Vars(r)

	firstName := params["firstName"]
	lastName := params["lastName"]

	person := Person{FirstName: firstName, LastName: lastName}

	person.SetHeroName()

	// serializedReturn, _ := json.Marshal(person)

	json.NewEncoder(w).Encode(person)
}

func Health(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /api/Heroes/Health getHealth
	//
	// Function used to check the health of the service
	//
	// ---
	// produces:
	// - text/plain
	// responses:
	//   '200':
	//     description: Health status

	fmt.Fprintf(w, "Service is healthy.")
}
