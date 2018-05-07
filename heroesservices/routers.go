package heroesservices

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route - object representing a route handler

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Route handler collection

type Routes []Route

// NewRouter - Constructor

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

// Index - Default route handler for service base uri

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Heroes Service")
}

var routes = Routes{
	Route{
		"Index",
		strings.ToUpper("Get"),
		"api/Heroes/Index",
		Index,
	},
	Route{
		"GetHero",
		strings.ToUpper("Get"),
		"/api/Heroes/Get/{firstName}/{lastName}",
		GetHero,
	},
}
