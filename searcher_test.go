package main

import (
	"strings"
	"testing"
)

const mockJSON = `{
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

func TestParseFlights(t *testing.T) {
	flights, err := parseFlights([]byte(mockJSON))
	if err != nil {
		t.Fatalf("parseFlights returned error: %v", err)
	}

	if len(flights) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(flights))
	}

	// Sorted by price ascending: 159, 189, 245
	if flights[0].Price != 159 {
		t.Errorf("flights[0].Price = %d, want 159", flights[0].Price)
	}
	if flights[1].Price != 189 {
		t.Errorf("flights[1].Price = %d, want 189", flights[1].Price)
	}
	if flights[2].Price != 245 {
		t.Errorf("flights[2].Price = %d, want 245", flights[2].Price)
	}

	// Cheapest is JetBlue B6 816
	if flights[0].Airline != "JetBlue" {
		t.Errorf("flights[0].Airline = %q, want \"JetBlue\"", flights[0].Airline)
	}
	if flights[0].FlightNumber != "B6 816" {
		t.Errorf("flights[0].FlightNumber = %q, want \"B6 816\"", flights[0].FlightNumber)
	}
	if flights[0].Duration != "5h 10m" {
		t.Errorf("flights[0].Duration = %q, want \"5h 10m\"", flights[0].Duration)
	}
}

func TestMissingSerpAPIKeyReturnsError(t *testing.T) {
	empty := ""
	apiKeyOverride = &empty
	defer func() { apiKeyOverride = nil }()

	_, err := searchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error for missing API key, got nil")
	}
	if !strings.Contains(err.Error(), "SERPAPI_KEY") {
		t.Errorf("expected error to contain \"SERPAPI_KEY\", got: %v", err)
	}
}
