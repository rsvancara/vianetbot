package util

import (
	"fmt"

	"net/http"
	"internal/internal/locationhelper"

	
)

// CtxKey Context Key
type CtxKey string

// ViaHTTPRequestResult ViaHttpRequestResult Object
type ViaHTTPRequestResult struct {
	
}

// ViaHTTPRequestQuery ViaHttpRequestQuery Object
type ViaHTTPRequestQuery struct {
	
}

//GeoIPContext get the geoip object
func GeoIPContext(r *http.Request) (locationhelper.GeoIP, error) {

	var geoIP locationhelper.GeoIP

	// Attempt to extract additional information from a context
	//var geoIP requestfilter.GeoIP
	var ctxKey CtxKey
	ctxKey = "geoip"

	if result := r.Context().Value(ctxKey); result != nil {

			fmt.Println("Found context")
			//fmt.Println(result)
			// Type Assertion....
			geoIP, ok := result.(locationhelper.GeoIP)
			if !ok {
					return geoIP, fmt.Errorf("could not perform type assertion on result to GeoIP type for ctxKey %s", ctxKey)
			}
			return geoIP, nil
	}
	return geoIP, fmt.Errorf("unable to find context for geoip %s", ctxKey)
}

// GetPageID get the page ID for a request
func GetPageID(r *http.Request) string {
	geoIP, err := GeoIPContext(r)
	if err != nil {
			fmt.Printf("error obtaining geoip context: %s\n", err)
	}

	return geoIP.PageID
}

