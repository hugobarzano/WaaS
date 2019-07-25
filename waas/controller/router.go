package controller

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"waas/mongo"
)


var controller = Controller{Storer: mongo.Storer{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route



var INSECURE_API = "/v0"

var routesV0 = Routes {
	Route {
		"Index",
		"GET",
		INSECURE_API,
		controller.IndexV0,
	},
	Route {
		"PushObject",
		"POST",
		INSECURE_API,
		controller.PushObjectV0,
	},
	}

var SECURE_API = "/v1"

var routesV1 = Routes {
	Route {
		"Index",
		"GET",
		SECURE_API,
		controller.IndexV1,
	},
	Route {
		"PushObject",
		"POST",
		SECURE_API,
		controller.PushObjectV1,
	},
}


// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routesV0 {
		var handler http.Handler
		log.Println(route.Pattern+":"+route.Name)
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range routesV1 {
		var handler http.Handler
		log.Println(route.Pattern+":"+route.Name)
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

