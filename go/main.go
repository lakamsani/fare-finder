// fare-finder: CLI tool to find the cheapest US domestic flight between two cities (Go port).
//
// Fixes https://github.com/lakamsani/fare-finder/issues/14
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	fare_finder "github.com/lakamsani/fare-finder/fare_finder"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Usage: fare-finder <city1> <state1> <city2> <state2>")
		fmt.Fprintln(os.Stderr, `Example: fare-finder "San Francisco" CA "New York" NY`)
		os.Exit(1)
	}

	city1 := titleCase(strings.TrimSpace(os.Args[1]))
	state1 := strings.ToUpper(strings.TrimSpace(os.Args[2]))
	city2 := titleCase(strings.TrimSpace(os.Args[3]))
	state2 := strings.ToUpper(strings.TrimSpace(os.Args[4]))

	origin, err := fare_finder.LookupAirport(city1, state1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	dest, err := fare_finder.LookupAirport(city2, state2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Printf("Searching for flights %s -> %s on %s...\n\n", origin, dest, tomorrow)

	flights, err := fare_finder.SearchFlights(origin, dest, tomorrow)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
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

// titleCase converts a string so that each word starts with an uppercase letter
// and the remaining letters are lowercase.
func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		runes := []rune(w)
		for j, r := range runes {
			if j == 0 {
				runes[j] = unicode.ToUpper(r)
			} else {
				runes[j] = unicode.ToLower(r)
			}
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
