package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"internal/internal/util"
	"internal/internal/views"
)

//GetRoutes get the routes for the application
func GetRoutes() *mux.Router {

	staticAssets, err := util.SiteTemplate("/static")
	if err != nil {
		log.Fatal(err.Error())
	}

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

	ServeStatic(r, "./"+staticAssets)
	http.Handle("/", r)

	return r

}

// ServeStatic  serve static content from the appropriate location
func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"css":     staticDirectory + "/css/",
		"images":  staticDirectory + "/images/",
		"scripts": staticDirectory + "/scripts/",
	}
	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
}
