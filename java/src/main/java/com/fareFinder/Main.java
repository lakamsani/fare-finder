package com.fareFinder;

import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.List;

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
            origin = Airport.lookup(city1, state1);
        } catch (IllegalArgumentException e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        try {
            dest = Airport.lookup(city2, state2);
        } catch (IllegalArgumentException e) {
            System.err.println("Error: " + e.getMessage());
            System.exit(1);
            return;
        }

        String tomorrow = LocalDate.now().plusDays(1).format(DateTimeFormatter.ofPattern("yyyy-MM-dd"));

        System.out.printf("Searching for flights %s \u2192 %s on %s...%n%n", origin, dest, tomorrow);

        List<Flight> flights;
        try {
            flights = FlightSearcher.search(origin, dest, tomorrow);
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
        System.out.printf("\uD83D\uDEEB Cheapest flight: %s \u2192 %s%n", origin, dest);
        System.out.printf("\u2708\uFE0F  %s %s%n", cheapest.airline(), cheapest.flightNumber());
        System.out.printf("\uD83D\uDCB0 $%d%n", cheapest.price());
        System.out.printf("\uD83D\uDD57 Departs %s | Arrives %s%n", cheapest.departureTime(), cheapest.arrivalTime());
        System.out.printf("\u23F1\uFE0F  Duration: %s%n", cheapest.duration());
    }

    static String titleCase(String s) {
        if (s == null || s.isEmpty()) return s;
        StringBuilder sb = new StringBuilder();
        boolean capitalizeNext = true;
        for (char c : s.toCharArray()) {
            if (Character.isWhitespace(c)) {
                capitalizeNext = true;
                sb.append(c);
            } else {
                if (capitalizeNext) {
                    sb.append(Character.toUpperCase(c));
                    capitalizeNext = false;
                } else {
                    sb.append(Character.toLowerCase(c));
                }
            }
        }
        return sb.toString();
    }
}
