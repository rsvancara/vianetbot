package locationhelper

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"bytes"

	"github.com/oschwald/geoip2-golang"
)

// GeoIP Object
type GeoIP struct {
	IsFound        bool
	IsPrivate      bool
	IPAddress      net.IP
	City           string
	CountryName    string
	CountryISOCode string
	TimeZone       string
	IsProxy        bool
	IsEU           bool
	PageID         string
}

// Search get geoip information from ipaddress
func (g *GeoIP) Search(ipaddress string) error {

	start := time.Now()

	ip := net.ParseIP(ipaddress)
	if ip == nil {
		g.IsFound = false
		return fmt.Errorf("error converting string to IP Address")
	}

	if _, err := os.Stat("db/GeoIP2-City.mmdb"); os.IsNotExist(err) {
		g.IsFound = false
		return fmt.Errorf("error opening city geodatabase")
	}

	db, err := geoip2.Open("db/GeoIP2-City.mmdb")
	if err != nil {
		g.IsFound = false
		return fmt.Errorf("error opening country geodatabase")
	}
	defer db.Close()

	record, err := db.City(ip)
	if err != nil {
		g.IsFound = false
		return fmt.Errorf("error getting database record: %s", err)
	}

	// Each language is represented in a map
	g.City = record.City.Names["en"]

	// Each language is represented in a map
	g.CountryName = record.Country.Names["en"]

	g.CountryISOCode = record.Country.IsoCode

	g.IPAddress = ip

	g.TimeZone = record.Location.TimeZone

	g.IsProxy = record.Traits.IsAnonymousProxy

	g.IsEU = record.Country.IsInEuropeanUnion

	elapsed := time.Since(start)
	log.Printf("geoipa took %s \n", elapsed)

	return nil
}

// IsPrivateSubnet - check to see if this ip is in a private subnet
func IsPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}


//ipRange - a structure that holds the start and end of a range of ip addresses
type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
	// TODO: Add IPV6 Ranges here
}
