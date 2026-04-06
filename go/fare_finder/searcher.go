package fare_finder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

// baseURL is the SerpAPI endpoint. It can be overridden in tests.
var baseURL = "https://serpapi.com/search"

// apiKeyOverride allows tests to inject a specific key.
// nil  = read from SERPAPI_KEY environment variable
// ""   = simulate missing key (triggers error)
// other = use this key directly
var apiKeyOverride *string

// APIError is returned when the SerpAPI responds with a non-200 status code.
type APIError struct {
	StatusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API returned status %d", e.StatusCode)
}

// SetBaseURL overrides the SerpAPI base URL. Intended for tests only.
func SetBaseURL(u string) {
	baseURL = u
}

// SetAPIKeyOverride sets the API key override used in tests.
// Pass nil to reset to environment variable resolution.
func SetAPIKeyOverride(key *string) {
	apiKeyOverride = key
}

// serpAPIResponse mirrors the relevant parts of the SerpAPI Google Flights JSON.
type serpAPIResponse struct {
	BestFlights  []flightGroup `json:"best_flights"`
	OtherFlights []flightGroup `json:"other_flights"`
}

type flightGroup struct {
	Flights []flightLeg `json:"flights"`
	Price   int         `json:"price"`
}

type flightLeg struct {
	Airline          string   `json:"airline"`
	FlightNumber     string   `json:"flight_number"`
	Duration         int      `json:"duration"`
	DepartureAirport timeNode `json:"departure_airport"`
	ArrivalAirport   timeNode `json:"arrival_airport"`
}

type timeNode struct {
	Time string `json:"time"`
}

// SearchFlights fetches flights from SerpAPI for the given origin, destination, and date.
//
// origin and dest must be 3-letter IATA codes.
// date must be in yyyy-MM-dd format.
// Returns flights sorted by price ascending.
func SearchFlights(origin, dest, date string) ([]Flight, error) {
	apiKey := resolveAPIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com")
	}

	params := url.Values{
		"engine":        {"google_flights"},
		"departure_id":  {origin},
		"arrival_id":    {dest},
		"outbound_date": {date},
		"currency":      {"USD"},
		"hl":            {"en"},
		"api_key":       {apiKey},
		"type":          {"2"},
	}
	reqURL := baseURL + "?" + params.Encode()

	resp, err := http.Get(reqURL) //nolint:noctx
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{StatusCode: resp.StatusCode}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return ParseFlights(string(body))
}

// ParseFlights parses a SerpAPI JSON response and returns flights sorted by price ascending.
func ParseFlights(jsonText string) ([]Flight, error) {
	var data serpAPIResponse
	if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	var flights []Flight

	for _, groups := range [][]flightGroup{data.BestFlights, data.OtherFlights} {
		for _, group := range groups {
			legs := group.Flights
			if len(legs) == 0 {
				continue
			}

			firstLeg := legs[0]
			lastLeg := legs[len(legs)-1]

			totalDuration := 0
			for _, leg := range legs {
				totalDuration += leg.Duration
			}

			flights = append(flights, Flight{
				Airline:       firstLeg.Airline,
				FlightNumber:  firstLeg.FlightNumber,
				Price:         group.Price,
				DepartureTime: firstLeg.DepartureAirport.Time,
				ArrivalTime:   lastLeg.ArrivalAirport.Time,
				Duration:      formatDuration(totalDuration),
			})
		}
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights, nil
}

// formatDuration converts total minutes into a human-readable "Xh Ym" string.
func formatDuration(minutes int) string {
	return fmt.Sprintf("%dh %dm", minutes/60, minutes%60)
}

// resolveAPIKey returns the API key to use, honouring the test override.
func resolveAPIKey() string {
	if apiKeyOverride != nil {
		return *apiKeyOverride
	}
	return os.Getenv("SERPAPI_KEY")
}
