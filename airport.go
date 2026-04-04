package main

import (
	"fmt"
	"strings"
)

var airports = map[string]string{
	"new york,ny":       "JFK",
	"los angeles,ca":    "LAX",
	"chicago,il":        "ORD",
	"houston,tx":        "IAH",
	"phoenix,az":        "PHX",
	"philadelphia,pa":   "PHL",
	"san antonio,tx":    "SAT",
	"san diego,ca":      "SAN",
	"dallas,tx":         "DFW",
	"san jose,ca":       "SJC",
	"austin,tx":         "AUS",
	"jacksonville,fl":   "JAX",
	"san francisco,ca":  "SFO",
	"columbus,oh":       "CMH",
	"charlotte,nc":      "CLT",
	"indianapolis,in":   "IND",
	"seattle,wa":        "SEA",
	"denver,co":         "DEN",
	"nashville,tn":      "BNA",
	"oklahoma city,ok":  "OKC",
	"el paso,tx":        "ELP",
	"washington,dc":     "DCA",
	"las vegas,nv":      "LAS",
	"louisville,ky":     "SDF",
	"baltimore,md":      "BWI",
	"milwaukee,wi":      "MKE",
	"albuquerque,nm":    "ABQ",
	"tucson,az":         "TUS",
	"fresno,ca":         "FAT",
	"sacramento,ca":     "SMF",
	"kansas city,mo":    "MCI",
	"atlanta,ga":        "ATL",
	"miami,fl":          "MIA",
	"minneapolis,mn":    "MSP",
	"portland,or":       "PDX",
	"detroit,mi":        "DTW",
	"boston,ma":         "BOS",
	"memphis,tn":        "MEM",
	"new orleans,la":    "MSY",
	"cleveland,oh":      "CLE",
	"tampa,fl":          "TPA",
	"orlando,fl":        "MCO",
}

// LookupAirport returns the IATA code for a given city and state.
// Returns an error if the city/state combination is not found.
func LookupAirport(city, state string) (string, error) {
	key := strings.ToLower(strings.TrimSpace(city)) + "," + strings.ToLower(strings.TrimSpace(state))
	code, ok := airports[key]
	if !ok {
		return "", fmt.Errorf("no airport found for %s, %s", city, state)
	}
	return code, nil
}
