// Package flights provides flight search and parsing via SerpAPI Google Flights.
package flights

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"time"
)

const serpAPIBaseURL = "https://serpapi.com/search"

// Flight holds details for a single flight result.
type Flight struct {
	Airline       string
	FlightNumber  string
	Price         int
	DepartureTime string
	ArrivalTime   string
	Duration      string
}

// Search queries SerpAPI for flights and returns them sorted by price (cheapest first).
// It reads SERPAPI_KEY from the environment.
func Search(origin, dest, date string) ([]Flight, error) {
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

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(serpAPIBaseURL + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SerpAPI returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return Parse(data), nil
}

// Parse converts a SerpAPI JSON response map into a price-sorted slice of Flights.
func Parse(data map[string]interface{}) []Flight {
	var results []Flight

	groups := concat(
		toSlice(data["best_flights"]),
		toSlice(data["other_flights"]),
	)

	for _, g := range groups {
		group, ok := g.(map[string]interface{})
		if !ok {
			continue
		}
		legs := toSlice(group["flights"])
		if len(legs) == 0 {
			continue
		}
		firstLeg, ok := legs[0].(map[string]interface{})
		if !ok {
			continue
		}
		lastLeg, ok := legs[len(legs)-1].(map[string]interface{})
		if !ok {
			continue
		}

		totalMinutes := 0
		for _, l := range legs {
			if leg, ok := l.(map[string]interface{}); ok {
				totalMinutes += int(toFloat(leg["duration"]))
			}
		}

		results = append(results, Flight{
			Airline:       toString(firstLeg["airline"]),
			FlightNumber:  toString(firstLeg["flight_number"]),
			Price:         int(toFloat(group["price"])),
			DepartureTime: toString(nested(firstLeg, "departure_airport", "time")),
			ArrivalTime:   toString(nested(lastLeg, "arrival_airport", "time")),
			Duration:      formatDuration(totalMinutes),
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Price < results[j].Price
	})
	return results
}

func formatDuration(minutes int) string {
	return fmt.Sprintf("%dh %dm", minutes/60, minutes%60)
}

// helpers

func toSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	s, _ := v.([]interface{})
	return s
}

func concat(a, b []interface{}) []interface{} {
	return append(a, b...)
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func toFloat(v interface{}) float64 {
	if v == nil {
		return 0
	}
	f, _ := v.(float64)
	return f
}

func nested(m map[string]interface{}, keys ...string) interface{} {
	var cur interface{} = m
	for _, k := range keys {
		mm, ok := cur.(map[string]interface{})
		if !ok {
			return nil
		}
		cur = mm[k]
	}
	return cur
}
