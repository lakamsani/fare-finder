// Package fare_finder provides types and logic for the fare-finder CLI (Go port).
package fare_finder

import "fmt"

// Flight represents a single flight option.
type Flight struct {
	Airline       string
	FlightNumber  string
	Price         int
	DepartureTime string
	ArrivalTime   string
	Duration      string
}

// String returns a human-readable representation of the flight.
func (f Flight) String() string {
	return fmt.Sprintf(
		"Flight{airline=%q, flightNumber=%q, price=%d, departureTime=%q, arrivalTime=%q, duration=%q}",
		f.Airline, f.FlightNumber, f.Price, f.DepartureTime, f.ArrivalTime, f.Duration,
	)
}
