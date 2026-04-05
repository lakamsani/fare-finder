package com.lakamsani.farefinder;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SearcherTest {

    private static final String MOCK_JSON = "{\n" +
            "    \"best_flights\": [\n" +
            "        {\n" +
            "            \"flights\": [\n" +
            "                {\n" +
            "                    \"departure_airport\": {\"time\": \"2024-01-15 08:05\"},\n" +
            "                    \"arrival_airport\": {\"time\": \"2024-01-15 16:23\"},\n" +
            "                    \"duration\": 318,\n" +
            "                    \"airline\": \"United Airlines\",\n" +
            "                    \"flight_number\": \"UA 101\"\n" +
            "                }\n" +
            "            ],\n" +
            "            \"price\": 189\n" +
            "        }\n" +
            "    ],\n" +
            "    \"other_flights\": [\n" +
            "        {\n" +
            "            \"flights\": [\n" +
            "                {\n" +
            "                    \"departure_airport\": {\"time\": \"2024-01-15 10:30\"},\n" +
            "                    \"arrival_airport\": {\"time\": \"2024-01-15 18:45\"},\n" +
            "                    \"duration\": 315,\n" +
            "                    \"airline\": \"Delta Air Lines\",\n" +
            "                    \"flight_number\": \"DL 405\"\n" +
            "                }\n" +
            "            ],\n" +
            "            \"price\": 245\n" +
            "        },\n" +
            "        {\n" +
            "            \"flights\": [\n" +
            "                {\n" +
            "                    \"departure_airport\": {\"time\": \"2024-01-15 06:00\"},\n" +
            "                    \"arrival_airport\": {\"time\": \"2024-01-15 14:10\"},\n" +
            "                    \"duration\": 310,\n" +
            "                    \"airline\": \"JetBlue\",\n" +
            "                    \"flight_number\": \"B6 816\"\n" +
            "                }\n" +
            "            ],\n" +
            "            \"price\": 159\n" +
            "        }\n" +
            "    ]\n" +
            "}";

    @AfterEach
    void resetOverride() {
        Searcher.apiKeyOverride = null;
    }

    @Test
    void testParseFlights() throws Exception {
        List<Flight> flights = Searcher.parseFlights(MOCK_JSON);

        assertEquals(3, flights.size(), "expected 3 flights");

        // Sorted by price ascending: 159, 189, 245
        assertEquals(159, flights.get(0).getPrice(), "flights[0].price");
        assertEquals(189, flights.get(1).getPrice(), "flights[1].price");
        assertEquals(245, flights.get(2).getPrice(), "flights[2].price");

        // Cheapest is JetBlue B6 816
        assertEquals("JetBlue", flights.get(0).getAirline(), "flights[0].airline");
        assertEquals("B6 816",  flights.get(0).getFlightNumber(), "flights[0].flightNumber");
        assertEquals("5h 10m",  flights.get(0).getDuration(), "flights[0].duration");
    }

    @Test
    void testMissingSerpAPIKeyReturnsError() {
        Searcher.apiKeyOverride = ""; // simulate missing key

        Exception ex = assertThrows(Exception.class,
                () -> Searcher.searchFlights("SFO", "JFK", "2024-01-15"));

        assertNotNull(ex.getMessage(), "error message should not be null");
        assertFalse(ex.getMessage().isEmpty(), "error message should not be empty");
        assertTrue(ex.getMessage().contains("SERPAPI_KEY"),
                "error message should mention SERPAPI_KEY, got: " + ex.getMessage());
    }
}
