package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()

	chiRoutes := testRoutes.(chi.Router) // casts to chi.Router

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, chiRoutes chi.Router, route string) {
	var found bool

	// walk through the routes looking for route
	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("Did not find %s in registered routes", route)
	}
}
