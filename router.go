package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes is all of the routes this server offers
type Routes []Route

// Route is used to represent an http endpoint
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// NewRouter creates and returns a new mux Router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}
	return router
}

var routes = Routes{
	Route{
		"Plex-Webhook",
		http.MethodPost,
		"/plex",
		Hook,
	},
}
