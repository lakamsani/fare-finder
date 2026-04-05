package com.lakamsani.farefinder;

import java.time.LocalDate;
import java.util.List;

/**
 * CLI entry point for fare-finder.
 */
public class Main {

    public static void main(String[] args) {
        if (args.length != 4) {
            System.err.println("Usage: fare-finder <city1> <state1> <city2> <state2>");
            System.err.println("Example: fare-finder \"San Francisco\" CA \"New York\" NY");
            System.exit(1);
        }

        String city1 = titleCase(args[0].trim());
        String state1 = args[1].trim().toUpperCase();
        String city2 = titleCase(args[2].trim());
        String state2 = args[3].trim().toUpperCase();

        String origin;
        String dest;
        try {
            origin = Airport.lookupAirport(city1, state1);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        try {
            dest = Airport.lookupAirport(city2, state2);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        String tomorrow = LocalDate.now().plusDays(1).toString();

        System.out.printf("Searching for flights %s -> %s on %s...%n%n", origin, dest, tomorrow);

        List<Flight> flights;
        try {
            flights = Searcher.searchFlights(origin, dest, tomorrow);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        if (flights.isEmpty()) {
            System.out.println("No flights found for this route and date. Try a different date or city pair.");
            System.exit(0);
        }

        Flight cheapest = flights.get(0);
        System.out.printf("Cheapest flight: %s -> %s%n", origin, dest);
        System.out.printf("%s %s%n", cheapest.getAirline(), cheapest.getFlightNumber());
        System.out.printf("$%d%n", cheapest.getPrice());
        System.out.printf("Departs %s | Arrives %s%n", cheapest.getDepartureTime(), cheapest.getArrivalTime());
        System.out.printf("Duration: %s%n", cheapest.getDuration());
    }

    /**
     * Converts a string so each word starts with an uppercase letter and remaining letters are lowercase.
     */
    static String titleCase(String s) {
        if (s == null || s.isEmpty()) {
            return s;
        }
        boolean capitalizeNext = true;
        StringBuilder sb = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (Character.isWhitespace(c)) {
                capitalizeNext = true;
                sb.append(c);
            } else if (capitalizeNext) {
                sb.append(Character.toUpperCase(c));
                capitalizeNext = false;
            } else {
                sb.append(Character.toLowerCase(c));
            }
        }
        return sb.toString();
    }
}
