package heroesservices

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	firstName := params["firstName"]
	lastName := params["lastName"]

	person := Person{FirstName: firstName, LastName: lastName}

	person.SetHeroName()

	serializedReturn, _ := json.Marshal(person)

	fmt.Fprintf(w, string(serializedReturn))
}
