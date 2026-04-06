package tests

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	fare_finder "github.com/lakamsani/fare-finder/fare_finder"
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

func TestParseFlights(t *testing.T) {
	flights, err := fare_finder.ParseFlights(mockJSON)
	if err != nil {
		t.Fatalf("ParseFlights returned unexpected error: %v", err)
	}

	if len(flights) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(flights))
	}

	// Sorted by price ascending: 159, 189, 245
	wantPrices := []int{159, 189, 245}
	for i, want := range wantPrices {
		if flights[i].Price != want {
			t.Errorf("flights[%d].Price = %d, want %d", i, flights[i].Price, want)
		}
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
	fare_finder.SetAPIKeyOverride(ptr("")) // simulate missing key
	defer fare_finder.SetAPIKeyOverride(nil)

	_, err := fare_finder.SearchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error for missing SERPAPI_KEY, got nil")
	}
	if msg := err.Error(); msg == "" {
		t.Error("error message should not be empty")
	}
	expected := "SERPAPI_KEY"
	if !containsString(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got: %q", expected, err.Error())
	}
}

func TestSearchFlightsAPIError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	fare_finder.SetBaseURL(ts.URL)
	fare_finder.SetAPIKeyOverride(ptr("test-key"))
	defer func() {
		fare_finder.SetBaseURL("https://serpapi.com/search")
		fare_finder.SetAPIKeyOverride(nil)
	}()

	_, err := fare_finder.SearchFlights("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error for non-200 status, got nil")
	}
	var target *fare_finder.APIError
	if !errors.As(err, &target) {
		// fall back to string check
		if !containsString(err.Error(), "500") {
			t.Errorf("expected error to mention status 500, got: %q", err.Error())
		}
	}
}

func TestSearchFlightsSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, mockJSON)
	}))
	defer ts.Close()

	fare_finder.SetBaseURL(ts.URL)
	fare_finder.SetAPIKeyOverride(ptr("test-key"))
	defer func() {
		fare_finder.SetBaseURL("https://serpapi.com/search")
		fare_finder.SetAPIKeyOverride(nil)
	}()

	flights, err := fare_finder.SearchFlights("SFO", "JFK", "2024-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(flights) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(flights))
	}
	if flights[0].Price != 159 {
		t.Errorf("cheapest flight price = %d, want 159", flights[0].Price)
	}
}

func ptr(s string) *string { return &s }

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
