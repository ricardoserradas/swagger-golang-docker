package heroesservices

import (
	"fmt"
	"net/http"
)

func GetHero(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing GetHero")
}
