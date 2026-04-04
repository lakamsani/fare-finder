package airports_test

import (
	"testing"

	"github.com/lakamsani/fare-finder/airports"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		city     string
		state    string
		expected string
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
		t.Run(tt.city+"-"+tt.state, func(t *testing.T) {
			got, err := airports.Lookup(tt.city, tt.state)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("Lookup(%q, %q) = %q; want %q", tt.city, tt.state, got, tt.expected)
			}
		})
	}
}

func TestLookupNotFound(t *testing.T) {
	_, err := airports.Lookup("Nonexistent City", "XX")
	if err == nil {
		t.Fatal("expected error for unknown city/state, got nil")
	}
}
