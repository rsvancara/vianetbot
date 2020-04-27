package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"internal/internal/views"
)

//GetRoutes get the routes for the application
func GetRoutes() *mux.Router {

	r := mux.NewRouter()

	// Index Page
	r.Handle(
		"/",
		handlers.LoggingHandler(
			os.Stdout,
			http.HandlerFunc(views.HomeHandler))).Methods("GET")

	// Test Page
	r.Handle(
		"/healthcheck957873",
		handlers.LoggingHandler(
			os.Stdout,
			http.HandlerFunc(views.HealthCheck))).Methods("GET")

	// Media Interface
	r.Handle(
		"/api/v1/risk",
		handlers.LoggingHandler(
			os.Stdout,
			http.HandlerFunc(views.RiskAPI))).Methods("GET", "POST")

	http.Handle("/", r)

	return r
}