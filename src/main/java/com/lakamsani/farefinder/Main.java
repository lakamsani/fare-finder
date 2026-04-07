package com.lakamsani.farefinder;

import java.time.LocalDate;
import java.util.List;

public class Main {
    static String titleCase(String s) {
        if (s == null || s.isEmpty()) return s;
        StringBuilder result = new StringBuilder();
        boolean capitalizeNext = true;
        for (char c : s.toCharArray()) {
            if (Character.isWhitespace(c)) {
                capitalizeNext = true;
                result.append(c);
            } else if (capitalizeNext) {
                result.append(Character.toUpperCase(c));
                capitalizeNext = false;
            } else {
                result.append(Character.toLowerCase(c));
            }
        }
        return result.toString();
    }

    public static void main(String[] args) {
        if (args.length != 4) {
            System.err.println("Usage: fare-finder <city1> <state1> <city2> <state2>");
            System.err.println("Example: fare-finder \"San Francisco\" CA \"New York\" NY");
            System.exit(1);
        }

        String city1 = titleCase(args[0].strip());
        String state1 = args[1].strip().toUpperCase();
        String city2 = titleCase(args[2].strip());
        String state2 = args[3].strip().toUpperCase();

        String origin;
        try {
            origin = Airport.lookupAirport(city1, state1);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        String dest;
        try {
            dest = Airport.lookupAirport(city2, state2);
        } catch (Exception e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        String tomorrow = LocalDate.now().plusDays(1).toString();

        System.out.println("Searching for flights " + origin + " -> " + dest + " on " + tomorrow + "...\n");

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
        System.out.println("Cheapest flight: " + origin + " -> " + dest);
        System.out.println(cheapest.getAirline() + " " + cheapest.getFlightNumber());
        System.out.println("$" + cheapest.getPrice());
        System.out.println("Departs " + cheapest.getDepartureTime() + " | Arrives " + cheapest.getArrivalTime());
        System.out.println("Duration: " + cheapest.getDuration());
    }
}
