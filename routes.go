package main

import (
	"net/http"
)

//Route . . .
//struct used to define a route,
//Name: description of the route
//Method: http method
//Pattern: actual route endpoint
//HandlerFunc: function to handle the route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice of routes that are registered with the mux router. All new routes can be defined here.
type Routes []Route

var routes = Routes{
	Route{
		"Register a new user in the system.", "POST", "/register", Register,
	},
	Route{
		"Login.", "POST", "/login", Login,
	},
}
