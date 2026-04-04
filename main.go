// fare-finder: find the cheapest US domestic flight between two cities.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lakamsani/fare-finder/airports"
	"github.com/lakamsani/fare-finder/flights"
)

func titleCase(s string) string {
	words := strings.Fields(strings.TrimSpace(s))
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "Usage: fare-finder <city1> <state1> <city2> <state2>\n")
		fmt.Fprintf(os.Stderr, "Example: fare-finder \"San Francisco\" CA \"New York\" NY\n")
		os.Exit(1)
	}

	city1 := titleCase(os.Args[1])
	state1 := strings.ToUpper(strings.TrimSpace(os.Args[2]))
	city2 := titleCase(os.Args[3])
	state2 := strings.ToUpper(strings.TrimSpace(os.Args[4]))

	origin, err := airports.Lookup(city1, state1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	dest, err := airports.Lookup(city2, state2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Printf("Searching for flights %s → %s on %s...\n\n", origin, dest, tomorrow)

	results, err := flights.Search(origin, dest, tomorrow)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(results) == 0 {
		fmt.Println("No flights found for this route and date. Try a different date or city pair.")
		os.Exit(0)
	}

	cheapest := results[0]
	fmt.Printf("🛫 Cheapest flight: %s → %s\n", origin, dest)
	fmt.Printf("✈️  %s %s\n", cheapest.Airline, cheapest.FlightNumber)
	fmt.Printf("💰 $%d\n", cheapest.Price)
	fmt.Printf("🕗 Departs %s | Arrives %s\n", cheapest.DepartureTime, cheapest.ArrivalTime)
	fmt.Printf("⏱️  Duration: %s\n", cheapest.Duration)
}
