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

public class Searcher {
    static String baseUrl = "https://serpapi.com/search";
    static String apiKeyOverride = null;

    public static List<Flight> searchFlights(String origin, String dest, String date) throws Exception {
        String apiKey;
        if (apiKeyOverride != null) {
            apiKey = apiKeyOverride;
        } else {
            apiKey = System.getenv("SERPAPI_KEY");
            if (apiKey == null) apiKey = "";
        }

        if (apiKey.isEmpty()) {
            throw new Exception("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com");
        }

        String url = baseUrl + "?engine=google_flights&departure_id=" + origin +
                     "&arrival_id=" + dest + "&outbound_date=" + date +
                     "&currency=USD&hl=en&api_key=" + apiKey + "&type=2";

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

    public static List<Flight> parseFlights(String json) throws Exception {
        ObjectMapper mapper = new ObjectMapper();
        JsonNode root = mapper.readTree(json);

        List<Flight> flights = new ArrayList<>();

        JsonNode bestFlights = root.path("best_flights");
        JsonNode otherFlights = root.path("other_flights");

        for (JsonNode group : concat(bestFlights, otherFlights)) {
            JsonNode legs = group.path("flights");
            if (!legs.isArray() || legs.size() == 0) continue;

            JsonNode first = legs.get(0);
            JsonNode last = legs.get(legs.size() - 1);

            int totalDuration = 0;
            for (JsonNode leg : legs) {
                totalDuration += leg.path("duration").asInt(0);
            }

            flights.add(new Flight(
                first.path("airline").asText(""),
                first.path("flight_number").asText(""),
                group.path("price").asInt(0),
                first.path("departure_airport").path("time").asText(""),
                last.path("arrival_airport").path("time").asText(""),
                formatDuration(totalDuration)
            ));
        }

        flights.sort(Comparator.comparingInt(Flight::getPrice));
        return flights;
    }

    private static List<JsonNode> concat(JsonNode a, JsonNode b) {
        List<JsonNode> result = new ArrayList<>();
        if (a.isArray()) for (JsonNode n : a) result.add(n);
        if (b.isArray()) for (JsonNode n : b) result.add(n);
        return result;
    }

    private static String formatDuration(int minutes) {
        return (minutes / 60) + "h " + (minutes % 60) + "m";
    }
}
