package views

import (
	//"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler Home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home")
}

// HealthCheck defines a healthcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "healthy")
}

// GetRequestAPI get the request information
func RiskAPI(w http.ResponseWriter, r *http.Request) {

	//errorMessage := "{\"status\":\"error\", \"message\": \"error: %s in %s\",\"file\":\"error\"}\n"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\":\"success\", \"message\": \"request recieved \"}\n")
	return
}
