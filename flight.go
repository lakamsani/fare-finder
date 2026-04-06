package main

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

func (f Flight) String() string {
	return fmt.Sprintf("Flight{airline='%s', flightNumber='%s', price=%d, departureTime='%s', arrivalTime='%s', duration='%s'}",
		f.Airline, f.FlightNumber, f.Price, f.DepartureTime, f.ArrivalTime, f.Duration)
}
