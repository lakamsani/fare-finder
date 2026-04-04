package com.fareFinder;

import com.google.gson.Gson;
import com.google.gson.JsonArray;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class FlightSearcher {
    static String BASE_URL = "https://serpapi.com/search";

    // Package-private for testing: empty string simulates missing key, null means use env var
    static String apiKeyOverride = null;

    public static List<Flight> search(String origin, String dest, String date) throws Exception {
        String apiKey = (apiKeyOverride != null) ? apiKeyOverride : System.getenv("SERPAPI_KEY");
        if (apiKey == null || apiKey.isEmpty()) {
            throw new IllegalStateException(
                "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com");
        }

        String url = String.format(
            "%s?engine=google_flights&departure_id=%s&arrival_id=%s&outbound_date=%s&currency=USD&hl=en&api_key=%s&type=2",
            BASE_URL, origin, dest, date, apiKey);

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(url))
                .GET()
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new RuntimeException("API returned status " + response.statusCode());
        }

        return parseFlights(response.body());
    }

    public static List<Flight> parseFlights(String json) {
        Gson gson = new Gson();
        JsonObject root = gson.fromJson(json, JsonObject.class);

        List<Flight> flights = new ArrayList<>();

        JsonArray bestFlights = root.has("best_flights") ? root.getAsJsonArray("best_flights") : new JsonArray();
        JsonArray otherFlights = root.has("other_flights") ? root.getAsJsonArray("other_flights") : new JsonArray();

        List<JsonElement> allGroups = new ArrayList<>();
        for (JsonElement e : bestFlights) allGroups.add(e);
        for (JsonElement e : otherFlights) allGroups.add(e);

        for (JsonElement groupElem : allGroups) {
            JsonObject group = groupElem.getAsJsonObject();
            JsonArray legs = group.getAsJsonArray("flights");
            if (legs == null || legs.size() == 0) continue;

            JsonObject firstLeg = legs.get(0).getAsJsonObject();
            JsonObject lastLeg = legs.get(legs.size() - 1).getAsJsonObject();

            int totalDuration = 0;
            for (JsonElement legElem : legs) {
                totalDuration += legElem.getAsJsonObject().get("duration").getAsInt();
            }

            String airline = firstLeg.get("airline").getAsString();
            String flightNumber = firstLeg.get("flight_number").getAsString();
            int price = group.get("price").getAsInt();
            String departureTime = firstLeg.getAsJsonObject("departure_airport").get("time").getAsString();
            String arrivalTime = lastLeg.getAsJsonObject("arrival_airport").get("time").getAsString();

            flights.add(new Flight(airline, flightNumber, price, departureTime, arrivalTime,
                formatDuration(totalDuration)));
        }

        flights.sort(Comparator.comparingInt(Flight::price));
        return flights;
    }

    static String formatDuration(int minutes) {
        return (minutes / 60) + "h " + (minutes % 60) + "m";
    }
}
