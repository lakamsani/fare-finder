package tests

import (
	"testing"

	fare_finder "github.com/lakamsani/fare-finder/fare_finder"
)

func TestKnownAirports(t *testing.T) {
	cases := []struct {
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

	for _, tc := range cases {
		got, err := fare_finder.LookupAirport(tc.city, tc.state)
		if err != nil {
			t.Errorf("LookupAirport(%q, %q): unexpected error: %v", tc.city, tc.state, err)
			continue
		}
		if got != tc.want {
			t.Errorf("LookupAirport(%q, %q) = %q, want %q", tc.city, tc.state, got, tc.want)
		}
	}
}

func TestUnknownCityReturnsError(t *testing.T) {
	_, err := fare_finder.LookupAirport("Nonexistent City", "XX")
	if err == nil {
		t.Fatal("expected error for unknown city, got nil")
	}
}
