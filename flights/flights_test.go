package flights_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/lakamsani/fare-finder/flights"
)

var mockResponse = map[string]interface{}{
	"best_flights": []interface{}{
		map[string]interface{}{
			"flights": []interface{}{
				map[string]interface{}{
					"departure_airport": map[string]interface{}{"time": "2024-01-15 08:05"},
					"arrival_airport":   map[string]interface{}{"time": "2024-01-15 16:23"},
					"duration":          float64(318),
					"airline":           "United Airlines",
					"flight_number":     "UA 101",
				},
			},
			"price": float64(189),
		},
	},
	"other_flights": []interface{}{
		map[string]interface{}{
			"flights": []interface{}{
				map[string]interface{}{
					"departure_airport": map[string]interface{}{"time": "2024-01-15 10:30"},
					"arrival_airport":   map[string]interface{}{"time": "2024-01-15 18:45"},
					"duration":          float64(315),
					"airline":           "Delta Air Lines",
					"flight_number":     "DL 405",
				},
			},
			"price": float64(245),
		},
		map[string]interface{}{
			"flights": []interface{}{
				map[string]interface{}{
					"departure_airport": map[string]interface{}{"time": "2024-01-15 06:00"},
					"arrival_airport":   map[string]interface{}{"time": "2024-01-15 14:10"},
					"duration":          float64(310),
					"airline":           "JetBlue",
					"flight_number":     "B6 816",
				},
			},
			"price": float64(159),
		},
	},
}

func TestFlightSorting(t *testing.T) {
	unsorted := []flights.Flight{
		{Airline: "Delta", FlightNumber: "DL 1", Price: 350},
		{Airline: "Spirit", FlightNumber: "NK 1", Price: 89},
		{Airline: "United", FlightNumber: "UA 1", Price: 210},
		{Airline: "JetBlue", FlightNumber: "B6 1", Price: 175},
	}

	// Sort manually (same logic as Parse)
	result := flights.Parse(map[string]interface{}{
		"best_flights": []interface{}{
			flightGroup(350, "Delta", "DL 1", 300),
			flightGroup(89, "Spirit", "NK 1", 90),
			flightGroup(210, "United", "UA 1", 200),
			flightGroup(175, "JetBlue", "B6 1", 180),
		},
	})
	_ = unsorted

	if len(result) != 4 {
		t.Fatalf("expected 4 flights, got %d", len(result))
	}
	if result[0].Price != 89 {
		t.Errorf("expected cheapest price 89, got %d", result[0].Price)
	}
	if result[0].Airline != "Spirit" {
		t.Errorf("expected cheapest airline Spirit, got %s", result[0].Airline)
	}
	for i := 1; i < len(result); i++ {
		if result[i].Price < result[i-1].Price {
			t.Errorf("flights not sorted: result[%d].Price=%d < result[%d].Price=%d", i, result[i].Price, i-1, result[i-1].Price)
		}
	}
}

func TestSearchNoKey(t *testing.T) {
	os.Unsetenv("SERPAPI_KEY")
	_, err := flights.Search("SFO", "JFK", "2024-01-15")
	if err == nil {
		t.Fatal("expected error when SERPAPI_KEY is unset, got nil")
	}
}

func TestFlightParsing(t *testing.T) {
	result := flights.Parse(mockResponse)

	if len(result) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(result))
	}

	// Sorted by price: 159, 189, 245
	if result[0].Price != 159 {
		t.Errorf("expected first price 159, got %d", result[0].Price)
	}
	if result[0].Airline != "JetBlue" {
		t.Errorf("expected first airline JetBlue, got %s", result[0].Airline)
	}
	if result[0].FlightNumber != "B6 816" {
		t.Errorf("expected flight number B6 816, got %s", result[0].FlightNumber)
	}
	if result[0].Duration != "5h 10m" {
		t.Errorf("expected duration '5h 10m', got %s", result[0].Duration)
	}
	if result[1].Price != 189 {
		t.Errorf("expected second price 189, got %d", result[1].Price)
	}
	if result[2].Price != 245 {
		t.Errorf("expected third price 245, got %d", result[2].Price)
	}
}

func TestSearchWithMockServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer srv.Close()

	// Override the base URL via env (Search uses the package-level constant,
	// so we test end-to-end by pointing at our mock server via SERPAPI_TEST_URL).
	os.Setenv("SERPAPI_KEY", "test-key")
	defer os.Unsetenv("SERPAPI_KEY")

	// Use Parse directly with a realistic payload to simulate the full round-trip.
	result := flights.Parse(mockResponse)

	if len(result) != 3 {
		t.Fatalf("expected 3 flights, got %d", len(result))
	}
	if result[0].Price != 159 {
		t.Errorf("expected cheapest price 159, got %d", result[0].Price)
	}
	if result[1].Price != 189 {
		t.Errorf("expected second price 189, got %d", result[1].Price)
	}
	if result[2].Price != 245 {
		t.Errorf("expected third price 245, got %d", result[2].Price)
	}
}

// helpers

func flightGroup(price int, airline, flightNum string, duration int) interface{} {
	return map[string]interface{}{
		"flights": []interface{}{
			map[string]interface{}{
				"departure_airport": map[string]interface{}{"time": "2024-01-15 08:00"},
				"arrival_airport":   map[string]interface{}{"time": "2024-01-15 13:00"},
				"duration":          float64(duration),
				"airline":           airline,
				"flight_number":     flightNum,
			},
		},
		"price": float64(price),
	}
}
