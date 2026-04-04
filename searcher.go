package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

// BaseURL is the SerpAPI endpoint. Overridable in tests.
var BaseURL = "https://serpapi.com/search"

// apiKeyOverride allows tests to inject a key (empty string = simulate missing key, nil = use env).
var apiKeyOverride *string

// SearchFlights fetches flights from SerpAPI for the given origin, destination, and date.
func SearchFlights(origin, dest, date string) ([]Flight, error) {
	var apiKey string
	if apiKeyOverride != nil {
		apiKey = *apiKeyOverride
	} else {
		apiKey = os.Getenv("SERPAPI_KEY")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com")
	}

	url := fmt.Sprintf(
		"%s?engine=google_flights&departure_id=%s&arrival_id=%s&outbound_date=%s&currency=USD&hl=en&api_key=%s&type=2",
		BaseURL, origin, dest, date, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return ParseFlights(string(body))
}

// serpResponse mirrors the relevant parts of the SerpAPI JSON response.
type serpResponse struct {
	BestFlights  []flightGroup `json:"best_flights"`
	OtherFlights []flightGroup `json:"other_flights"`
}

type flightGroup struct {
	Flights []legInfo `json:"flights"`
	Price   int       `json:"price"`
}

type legInfo struct {
	DepartureAirport airportInfo `json:"departure_airport"`
	ArrivalAirport   airportInfo `json:"arrival_airport"`
	Duration         int         `json:"duration"`
	Airline          string      `json:"airline"`
	FlightNumber     string      `json:"flight_number"`
}

type airportInfo struct {
	Time string `json:"time"`
}

// ParseFlights parses a SerpAPI JSON response and returns flights sorted by price ascending.
func ParseFlights(jsonStr string) ([]Flight, error) {
	var resp serpResponse
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	allGroups := append(resp.BestFlights, resp.OtherFlights...)
	var flights []Flight

	for _, group := range allGroups {
		if len(group.Flights) == 0 {
			continue
		}

		firstLeg := group.Flights[0]
		lastLeg := group.Flights[len(group.Flights)-1]

		totalDuration := 0
		for _, leg := range group.Flights {
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

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights, nil
}

// formatDuration converts total minutes into a human-readable "Xh Ym" string.
func formatDuration(minutes int) string {
	return fmt.Sprintf("%dh %dm", minutes/60, minutes%60)
}
