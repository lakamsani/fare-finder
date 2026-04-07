package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const mockJSON = `{
    "best_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 08:05"},
                    "arrival_airport":   {"time": "2024-01-15 16:23"},
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
                    "arrival_airport":   {"time": "2024-01-15 18:45"},
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
                    "arrival_airport":   {"time": "2024-01-15 14:10"},
                    "duration": 310,
                    "airline": "JetBlue",
                    "flight_number": "B6 816"
                }
            ],
            "price": 159
        }
    ]
}`

func TestParseFlights_ValidJSON(t *testing.T) {
	flights, err := ParseFlights(mockJSON)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
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

func TestParseFlights_EmptyJSON(t *testing.T) {
	flights, err := ParseFlights(`{}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(flights) != 0 {
		t.Errorf("expected 0 flights, got %d", len(flights))
	}
}

func TestParseFlights_InvalidJSON(t *testing.T) {
	_, err := ParseFlights(`not valid json`)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestSearchFlights_MissingSerpAPIKey(t *testing.T) {
	old := os.Getenv("SERPAPI_KEY")
	os.Unsetenv("SERPAPI_KEY")
	defer os.Setenv("SERPAPI_KEY", old)

	_, err := SearchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error for missing SERPAPI_KEY, got nil")
	}

	msg := err.Error()
	if msg == "" {
		t.Error("error message should not be empty")
	}

	found := false
	for i := 0; i+len("SERPAPI_KEY") <= len(msg); i++ {
		if msg[i:i+len("SERPAPI_KEY")] == "SERPAPI_KEY" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'SERPAPI_KEY' in error message, got: %q", msg)
	}
}

func TestSearchFlights_HTTPTest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockJSON)) //nolint:errcheck
	}))
	defer srv.Close()

	// Override base URL and set a fake API key.
	origURL := baseURL
	baseURL = srv.URL
	defer func() { baseURL = origURL }()

	os.Setenv("SERPAPI_KEY", "test-key")
	defer os.Unsetenv("SERPAPI_KEY")

	flights, err := SearchFlights("SFO", "JFK", "2024-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(flights) != 3 {
		t.Errorf("expected 3 flights, got %d", len(flights))
	}
}
