package com.lakamsani.farefinder;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

import java.util.List;

public class SearcherTest {

    private static final String MOCK_JSON = """
            {
                "best_flights": [
                    {
                        "flights": [
                            {
                                "departure_airport": {"time": "2024-01-15 08:05"},
                                "arrival_airport": {"time": "2024-01-15 16:23"},
                                "duration": 318,
                                "airline": "United Airlines",
                                "flight_number": "UA 101"
                            }
                        ],
                        "price": 189
                    }
                ],
                "other_flights": [
                    {
                        "flights": [
                            {
                                "departure_airport": {"time": "2024-01-15 10:30"},
                                "arrival_airport": {"time": "2024-01-15 18:45"},
                                "duration": 315,
                                "airline": "Delta Air Lines",
                                "flight_number": "DL 405"
                            }
                        ],
                        "price": 245
                    },
                    {
                        "flights": [
                            {
                                "departure_airport": {"time": "2024-01-15 06:00"},
                                "arrival_airport": {"time": "2024-01-15 14:10"},
                                "duration": 310,
                                "airline": "JetBlue",
                                "flight_number": "B6 816"
                            }
                        ],
                        "price": 159
                    }
                ]
            }
            """;

    @Test
    public void testParseFlights() throws Exception {
        List<Flight> flights = Searcher.parseFlights(MOCK_JSON);

        assertEquals(3, flights.size());

        // Sorted by price ascending: 159, 189, 245
        assertEquals(159, flights.get(0).getPrice());
        assertEquals(189, flights.get(1).getPrice());
        assertEquals(245, flights.get(2).getPrice());

        // Cheapest is JetBlue B6 816
        assertEquals("JetBlue", flights.get(0).getAirline());
        assertEquals("B6 816", flights.get(0).getFlightNumber());
        assertEquals("5h 10m", flights.get(0).getDuration());
    }

    @Test
    public void testMissingSerpAPIKeyReturnsError() {
        Searcher.apiKeyOverride = "";
        try {
            Exception ex = assertThrows(Exception.class,
                    () -> Searcher.searchFlights("SFO", "JFK", "2024-01-15"));
            assertTrue(ex.getMessage().contains("SERPAPI_KEY"),
                    "Expected message to contain SERPAPI_KEY but got: " + ex.getMessage());
        } finally {
            Searcher.apiKeyOverride = null;
        }
    }
}
