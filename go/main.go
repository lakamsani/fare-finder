package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func main() {
	args := os.Args[1:]

	if len(args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage: fare-finder <city1> <state1> <city2> <state2>")
		fmt.Fprintln(os.Stderr, `Example: fare-finder "San Francisco" CA "New York" NY`)
		os.Exit(1)
	}

	city1 := titleCase(strings.TrimSpace(args[0]))
	state1 := strings.ToUpper(strings.TrimSpace(args[1]))
	city2 := titleCase(strings.TrimSpace(args[2]))
	state2 := strings.ToUpper(strings.TrimSpace(args[3]))

	origin, err := LookupAirport(city1, state1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	dest, err := LookupAirport(city2, state2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Printf("Searching for flights %s -> %s on %s...\n\n", origin, dest, tomorrow)

	flights, err := SearchFlights(origin, dest, tomorrow)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	if len(flights) == 0 {
		fmt.Println("No flights found for this route and date. Try a different date or city pair.")
		os.Exit(0)
	}

	cheapest := flights[0]
	fmt.Printf("Cheapest flight: %s -> %s\n", origin, dest)
	fmt.Printf("%s %s\n", cheapest.Airline, cheapest.FlightNumber)
	fmt.Printf("$%d\n", cheapest.Price)
	fmt.Printf("Departs %s | Arrives %s\n", cheapest.DepartureTime, cheapest.ArrivalTime)
	fmt.Printf("Duration: %s\n", cheapest.Duration)
}
