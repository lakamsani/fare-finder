package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func titleCase(s string) string {
	if s == "" {
		return s
	}
	capitalizeNext := true
	var sb strings.Builder
	for _, c := range s {
		if unicode.IsSpace(c) {
			capitalizeNext = true
			sb.WriteRune(c)
		} else if capitalizeNext {
			sb.WriteRune(unicode.ToUpper(c))
			capitalizeNext = false
		} else {
			sb.WriteRune(unicode.ToLower(c))
		}
	}
	return sb.String()
}

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Usage: fare-finder <city1> <state1> <city2> <state2>")
		fmt.Fprintln(os.Stderr, "Example: fare-finder \"San Francisco\" CA \"New York\" NY")
		os.Exit(1)
	}

	city1 := titleCase(strings.TrimSpace(os.Args[1]))
	state1 := strings.ToUpper(strings.TrimSpace(os.Args[2]))
	city2 := titleCase(strings.TrimSpace(os.Args[3]))
	state2 := strings.ToUpper(strings.TrimSpace(os.Args[4]))

	origin, err := lookupAirport(city1, state1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	dest, err := lookupAirport(city2, state2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Printf("Searching for flights %s -> %s on %s...\n\n", origin, dest, tomorrow)

	flights, err := searchFlights(origin, dest, tomorrow)
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
