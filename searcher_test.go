package main

import (
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
	flights, err := ParseFlights(mockJSON)
	if err != nil {
		t.Fatalf("ParseFlights returned error: %v", err)
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
		t.Errorf("flights[0].Airline = %q, want %q", flights[0].Airline, "JetBlue")
	}
	if flights[0].FlightNumber != "B6 816" {
		t.Errorf("flights[0].FlightNumber = %q, want %q", flights[0].FlightNumber, "B6 816")
	}
	if flights[0].Duration != "5h 10m" {
		t.Errorf("flights[0].Duration = %q, want %q", flights[0].Duration, "5h 10m")
	}
}

func TestMissingSerpAPIKeyReturnsError(t *testing.T) {
	empty := ""
	apiKeyOverride = &empty
	defer func() { apiKeyOverride = nil }()

	_, err := SearchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error for missing SERPAPI_KEY, got nil")
	}
	if got := err.Error(); len(got) == 0 {
		t.Error("expected non-empty error message")
	}
	// Should mention SERPAPI_KEY
	if !contains(err.Error(), "SERPAPI_KEY") {
		t.Errorf("error message %q does not mention SERPAPI_KEY", err.Error())
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsStr(s, substr))
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
