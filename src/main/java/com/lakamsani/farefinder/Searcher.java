package com.lakamsani.farefinder;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/**
 * Handles flight search via SerpAPI and JSON parsing.
 */
public class Searcher {

    /** SerpAPI endpoint. Package-private so tests can override. */
    static String baseUrl = "https://serpapi.com/search";

    /**
     * API key override. Package-private.
     * null  = read from SERPAPI_KEY environment variable
     * ""    = simulate missing key (triggers error)
     * other = use this key directly
     */
    static String apiKeyOverride = null;

    private static final ObjectMapper mapper = new ObjectMapper();

    /**
     * Fetches flights from SerpAPI for the given origin, destination, and date.
     *
     * @param origin 3-letter IATA origin code
     * @param dest   3-letter IATA destination code
     * @param date   departure date in yyyy-MM-dd format
     * @return list of flights sorted by price ascending
     * @throws Exception if the API key is missing, the HTTP request fails, or JSON parsing fails
     */
    public static List<Flight> searchFlights(String origin, String dest, String date) throws Exception {
        String apiKey;
        if (apiKeyOverride != null) {
            apiKey = apiKeyOverride;
        } else {
            apiKey = System.getenv("SERPAPI_KEY");
            if (apiKey == null) {
                apiKey = "";
            }
        }

        if (apiKey.isEmpty()) {
            throw new Exception("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com");
        }

        String url = String.format(
                "%s?engine=google_flights&departure_id=%s&arrival_id=%s&outbound_date=%s&currency=USD&hl=en&api_key=%s&type=2",
                baseUrl, origin, dest, date, apiKey
        );

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(url))
                .GET()
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new Exception("API returned status " + response.statusCode());
        }

        return parseFlights(response.body());
    }

    /**
     * Parses a SerpAPI JSON response and returns flights sorted by price ascending.
     *
     * @param json the raw JSON string from SerpAPI
     * @return list of flights sorted by price ascending
     * @throws Exception if JSON parsing fails
     */
    public static List<Flight> parseFlights(String json) throws Exception {
        JsonNode root;
        try {
            root = mapper.readTree(json);
        } catch (Exception e) {
            throw new Exception("failed to parse JSON: " + e.getMessage(), e);
        }

        List<Flight> flights = new ArrayList<>();

        // Process best_flights and other_flights
        String[] groupKeys = {"best_flights", "other_flights"};
        for (String groupKey : groupKeys) {
            JsonNode groups = root.path(groupKey);
            if (groups.isMissingNode() || !groups.isArray()) {
                continue;
            }
            for (JsonNode group : groups) {
                JsonNode flightLegs = group.path("flights");
                if (!flightLegs.isArray() || flightLegs.size() == 0) {
                    continue;
                }
                int price = group.path("price").asInt(0);

                JsonNode firstLeg = flightLegs.get(0);
                JsonNode lastLeg = flightLegs.get(flightLegs.size() - 1);

                String airline = firstLeg.path("airline").asText("");
                String flightNumber = firstLeg.path("flight_number").asText("");
                String departureTime = firstLeg.path("departure_airport").path("time").asText("");
                String arrivalTime = lastLeg.path("arrival_airport").path("time").asText("");

                int totalDuration = 0;
                for (JsonNode leg : flightLegs) {
                    totalDuration += leg.path("duration").asInt(0);
                }

                flights.add(new Flight(airline, flightNumber, price, departureTime, arrivalTime,
                        formatDuration(totalDuration)));
            }
        }

        flights.sort(Comparator.comparingInt(Flight::getPrice));
        return flights;
    }

    /**
     * Converts total minutes into a human-readable "Xh Ym" string.
     */
    private static String formatDuration(int minutes) {
        return String.format("%dh %dm", minutes / 60, minutes % 60);
    }
}
