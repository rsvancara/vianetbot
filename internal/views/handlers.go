package views

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"encoding/json"
	"internal/internal/locationhelper"
	"net"
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

// RiskAPI get the request information
func RiskAPI(w http.ResponseWriter, r *http.Request) {

  type Req struct {
		Name string `json:"name"`
		ServerName string `json:"server_name"`
		RequestID string `json:"request_id"`
		QueryString string `json:"query_string"`
		Request string `json:"request"`
		RemoteAddr string `json:"remote_addr"`
		RequestLength string `json:"request_length"`
		RequestURI string `json:"request_length"`
		RequestScheme string `json:"request_scheme"`
	  SSLCipher string `json:"ssl_cipher"`
		MSec string `json:"msec"`
		UserAgent string `json:"user_agent"`
	}

	var d Req

	errorMessage := "{\"status\":\"error\", \"message\": \"request failed with error %s \"}\n"
	
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, errorMessage, err)
					return
	}

	fmt.Println(d)

	// Added for testing purposes
	if locationhelper.IsPrivateSubnet(net.ParseIP(d.RemoteAddr)) {
		d.RemoteAddr = "200.23.143.45"
	}

	var geoIP locationhelper.GeoIP

	err = geoIP.Search(d.RemoteAddr)
  if err != nil {
    fmt.Printf("Error IP Address not found in the database for IP Address: %s with error %s\n", d.RemoteAddr, err)
	}
	
	fmt.Println(geoIP)

	//errorMessage := "{\"status\":\"error\", \"message\": \"error: %s in %s\",\"file\":\"error\"}\n"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\":\"success\", \"message\": \"request recieved \"}\n")



}
