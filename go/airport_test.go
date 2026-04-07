package main

import (
	"testing"
)

func TestLookupAirport_KnownCities(t *testing.T) {
	cases := []struct {
		city, state, expected string
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
		t.Run(tc.city+","+tc.state, func(t *testing.T) {
			got, err := LookupAirport(tc.city, tc.state)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("LookupAirport(%q, %q) = %q, want %q", tc.city, tc.state, got, tc.expected)
			}
		})
	}
}

func TestLookupAirport_UnknownCity(t *testing.T) {
	_, err := LookupAirport("Nonexistent City", "XX")
	if err == nil {
		t.Fatal("expected error for unknown city, got nil")
	}
	if err.Error() == "" {
		t.Error("error message should not be empty")
	}
}
