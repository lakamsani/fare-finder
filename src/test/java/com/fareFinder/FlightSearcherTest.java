package com.fareFinder;

import org.junit.After;
import org.junit.Test;
import static org.junit.Assert.*;

import java.util.List;

public class FlightSearcherTest {

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

    @After
    public void resetOverrides() {
        FlightSearcher.apiKeyOverride = null;
        FlightSearcher.BASE_URL = "https://serpapi.com/search";
    }

    @Test
    public void testParseFlights() {
        List<Flight> flights = FlightSearcher.parseFlights(MOCK_JSON);

        assertEquals(3, flights.size());

        // Sorted by price ascending: 159, 189, 245
        assertEquals(159, flights.get(0).price());
        assertEquals(189, flights.get(1).price());
        assertEquals(245, flights.get(2).price());

        // First flight is JetBlue B6 816
        assertEquals("JetBlue", flights.get(0).airline());
        assertEquals("B6 816", flights.get(0).flightNumber());
        assertEquals("5h 10m", flights.get(0).duration());
    }

    @Test
    public void testMissingSerpApiKeyThrows() {
        FlightSearcher.apiKeyOverride = "";
        try {
            FlightSearcher.search("SFO", "JFK", "2024-01-15");
            fail("Expected exception for missing SERPAPI_KEY");
        } catch (Exception e) {
            assertTrue(e.getMessage().contains("SERPAPI_KEY"));
        }
    }
}
