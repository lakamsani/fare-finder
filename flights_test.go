package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"testing"
)

func TestAirportLookup(t *testing.T) {
	tests := []struct {
		city, state, want string
	}{
		{"San Francisco", "CA", "SFO"},
		{"New York", "NY", "JFK"},
		{"Chicago", "IL", "ORD"},
		{"Atlanta", "GA", "ATL"},
		{"Denver", "CO", "DEN"},
		{"Seattle", "WA", "SEA"},
		{"Miami", "FL", "MIA"},
		{"Las Vegas", "NV", "LAS"},
		{"Boston", "MA", "BOS"},
		{"Dallas", "TX", "DFW"},
	}

	for _, tt := range tests {
		got, err := LookupAirport(tt.city, tt.state)
		if err != nil {
			t.Errorf("LookupAirport(%q, %q) returned error: %v", tt.city, tt.state, err)
			continue
		}
		if got != tt.want {
			t.Errorf("LookupAirport(%q, %q) = %q, want %q", tt.city, tt.state, got, tt.want)
		}
	}
}

func TestAirportLookupNotFound(t *testing.T) {
	_, err := LookupAirport("Nonexistent City", "XX")
	if err == nil {
		t.Error("expected error for unknown city, got nil")
	}
}

func TestFlightSorting(t *testing.T) {
	flights := []Flight{
		{Airline: "Delta", Price: 350},
		{Airline: "Spirit", Price: 89},
		{Airline: "United", Price: 210},
		{Airline: "JetBlue", Price: 175},
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].Price < flights[j].Price
	})

	if flights[0].Price != 89 {
		t.Errorf("expected cheapest flight price 89, got %d", flights[0].Price)
	}
	if flights[0].Airline != "Spirit" {
		t.Errorf("expected cheapest airline Spirit, got %s", flights[0].Airline)
	}
	for i := 1; i < len(flights); i++ {
		if flights[i].Price < flights[i-1].Price {
			t.Errorf("flights not sorted: index %d price %d < index %d price %d", i, flights[i].Price, i-1, flights[i-1].Price)
		}
	}
}

func TestSearchFlightsNoKey(t *testing.T) {
	os.Unsetenv("SERPAPI_KEY")
	_, err := SearchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Error("expected error when SERPAPI_KEY is not set, got nil")
	}
}

func TestFlightParsing(t *testing.T) {
	mockResponse := `{
		"best_flights": [
			{
				"flights": [
					{
						"departure_airport": {"time": "2024-01-15 08:05"},
						"arrival_airport": {"time": "2024-01-15 16:23"},
						"duration": 318,
						"airline": "United Airlines",
						"flight_number": "UA 101"
					}
				],
				"price": 189
			}
		],
		"other_flights": [
			{
				"flights": [
					{
						"departure_airport": {"time": "2024-01-15 10:30"},
						"arrival_airport": {"time": "2024-01-15 18:45"},
						"duration": 315,
						"airline": "Delta Air Lines",
						"flight_number": "DL 405"
					}
				],
				"price": 245
			},
			{
				"flights": [
					{
						"departure_airport": {"time": "2024-01-15 06:00"},
						"arrival_airport": {"time": "2024-01-15 14:10"},
						"duration": 310,
						"airline": "JetBlue",
						"flight_number": "B6 816"
					}
				],
				"price": 159
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	origBaseURL := serpAPIBaseURL
	serpAPIBaseURL = server.URL
	defer func() { serpAPIBaseURL = origBaseURL }()

	os.Setenv("SERPAPI_KEY", "test-key")
	defer os.Unsetenv("SERPAPI_KEY")

	flights, err := SearchFlights("SFO", "JFK", "2024-01-15")
	if err != nil {
		t.Fatalf("SearchFlights returned error: %v", err)
	}

	if len(flights) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(flights))
	}

	// Should be sorted by price: 159, 189, 245
	if flights[0].Price != 159 {
		t.Errorf("expected cheapest flight price 159, got %d", flights[0].Price)
	}
	if flights[0].Airline != "JetBlue" {
		t.Errorf("expected cheapest airline JetBlue, got %s", flights[0].Airline)
	}
	if flights[1].Price != 189 {
		t.Errorf("expected second flight price 189, got %d", flights[1].Price)
	}
	if flights[2].Price != 245 {
		t.Errorf("expected third flight price 245, got %d", flights[2].Price)
	}
	if flights[0].FlightNumber != "B6 816" {
		t.Errorf("expected flight number B6 816, got %s", flights[0].FlightNumber)
	}
	if flights[0].Duration != "5h 10m" {
		t.Errorf("expected duration 5h 10m, got %s", flights[0].Duration)
	}
}
