package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

// baseURL is the SerpAPI endpoint. Override in tests via the package-level var.
var baseURL = "https://serpapi.com/search"

// SearchFlights fetches flights from SerpAPI for the given origin, destination, and date.
// Returns flights sorted by price ascending.
func SearchFlights(origin, dest, date string) ([]Flight, error) {
	apiKey := os.Getenv("SERPAPI_KEY")
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
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return ParseFlights(string(body))
}

// serpAPIResponse mirrors the relevant parts of a SerpAPI Google Flights response.
type serpAPIResponse struct {
	BestFlights  []flightGroup `json:"best_flights"`
	OtherFlights []flightGroup `json:"other_flights"`
}

type flightGroup struct {
	Flights []legData `json:"flights"`
	Price   int       `json:"price"`
}

type legData struct {
	DepartureAirport airportTime `json:"departure_airport"`
	ArrivalAirport   airportTime `json:"arrival_airport"`
	Duration         int         `json:"duration"`
	Airline          string      `json:"airline"`
	FlightNumber     string      `json:"flight_number"`
}

type airportTime struct {
	Time string `json:"time"`
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
