package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

var baseURL = "https://serpapi.com/search"
var apiKeyOverride *string

func searchFlights(origin, dest, date string) ([]Flight, error) {
	var apiKey string
	if apiKeyOverride != nil {
		apiKey = *apiKeyOverride
	} else {
		apiKey = os.Getenv("SERPAPI_KEY")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com")
	}

	url := fmt.Sprintf("%s?engine=google_flights&departure_id=%s&arrival_id=%s&outbound_date=%s&currency=USD&hl=en&api_key=%s&type=2",
		baseURL, origin, dest, date, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseFlights(body)
}

func parseFlights(data []byte) ([]Flight, error) {
	var root struct {
		BestFlights  []flightGroup `json:"best_flights"`
		OtherFlights []flightGroup `json:"other_flights"`
	}
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	var flights []Flight
	for _, group := range append(root.BestFlights, root.OtherFlights...) {
		if len(group.Flights) == 0 {
			continue
		}
		first := group.Flights[0]
		last := group.Flights[len(group.Flights)-1]

		totalDuration := 0
		for _, leg := range group.Flights {
			totalDuration += leg.Duration
		}

		flights = append(flights, Flight{
			Airline:       first.Airline,
			FlightNumber:  first.FlightNumber,
			Price:         group.Price,
			DepartureTime: first.DepartureAirport.Time,
			ArrivalTime:   last.ArrivalAirport.Time,
			Duration:      formatDuration(totalDuration),
		})
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	return flights, nil
}

type flightGroup struct {
	Flights []legInfo `json:"flights"`
	Price   int       `json:"price"`
}

type legInfo struct {
	DepartureAirport struct {
		Time string `json:"time"`
	} `json:"departure_airport"`
	ArrivalAirport struct {
		Time string `json:"time"`
	} `json:"arrival_airport"`
	Duration     int    `json:"duration"`
	Airline      string `json:"airline"`
	FlightNumber string `json:"flight_number"`
}

func formatDuration(minutes int) string {
	return fmt.Sprintf("%dh %dm", minutes/60, minutes%60)
}
