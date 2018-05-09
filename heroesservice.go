package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func SwaggerDoc(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /doc/{fileName} swaggerDoc
	//
	// Use swagger.json as a file name to return the OpenAPI swagger documentation in JSON.
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: fileName
	//   in: path
	//   description: use swagger.json
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: Swagger documentation

	params := mux.Vars(r)

	fileName := params["fileName"]

	jsonFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "%s\n", string(jsonFile))
}
