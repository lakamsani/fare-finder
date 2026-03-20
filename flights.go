package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

// Flight represents a single flight result.
type Flight struct {
	Airline       string
	FlightNumber  string
	Price         int
	DepartureTime string
	ArrivalTime   string
	Duration      string
}

// serpAPIResponse is the top-level SerpAPI JSON structure.
type serpAPIResponse struct {
	BestFlights  []flightGroup `json:"best_flights"`
	OtherFlights []flightGroup `json:"other_flights"`
}

type flightGroup struct {
	Flights []flightLeg `json:"flights"`
	Price   int         `json:"price"`
}

type flightLeg struct {
	DepartureAirport airportInfo `json:"departure_airport"`
	ArrivalAirport   airportInfo `json:"arrival_airport"`
	Duration         int         `json:"duration"`
	Airline          string      `json:"airline"`
	FlightNumber     string      `json:"flight_number"`
}

type airportInfo struct {
	Time string `json:"time"`
}

// serpAPIBaseURL can be overridden in tests.
var serpAPIBaseURL = "https://serpapi.com/search"

// SearchFlights queries SerpAPI for flights and returns them sorted by price (cheapest first).
func SearchFlights(origin, dest, date string) ([]Flight, error) {
	apiKey := os.Getenv("SERPAPI_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com")
	}

	url := fmt.Sprintf("%s?engine=google_flights&departure_id=%s&arrival_id=%s&outbound_date=%s&currency=USD&hl=en&api_key=%s&type=2",
		serpAPIBaseURL, origin, dest, date, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return ParseFlights(body)
}

// ParseFlights parses the SerpAPI JSON response into a sorted slice of flights.
func ParseFlights(data []byte) ([]Flight, error) {
	var result serpAPIResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %w", err)
	}

	var flights []Flight

	allGroups := append(result.BestFlights, result.OtherFlights...)
	for _, group := range allGroups {
		if len(group.Flights) == 0 {
			continue
		}
		leg := group.Flights[0]
		totalDuration := 0
		for _, l := range group.Flights {
			totalDuration += l.Duration
		}
		flights = append(flights, Flight{
			Airline:       leg.Airline,
			FlightNumber:  leg.FlightNumber,
			Price:         group.Price,
			DepartureTime: leg.DepartureAirport.Time,
			ArrivalTime:   group.Flights[len(group.Flights)-1].ArrivalAirport.Time,
			Duration:      formatDuration(totalDuration),
		})
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights, nil
}

func formatDuration(minutes int) string {
	return fmt.Sprintf("%dh %dm", minutes/60, minutes%60)
}
