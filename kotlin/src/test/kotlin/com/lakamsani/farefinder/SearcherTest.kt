package com.lakamsani.farefinder

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertFailsWith
import kotlin.test.assertFalse
import kotlin.test.assertNotNull
import kotlin.test.assertTrue

class SearcherTest {

    companion object {
        private val MOCK_JSON = """
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
        """.trimIndent()
    }

    @org.junit.jupiter.api.AfterEach
    fun resetOverride() {
        Searcher.apiKeyOverride = null
    }

    @Test
    fun testParseFlights() {
        val flights = Searcher.parseFlights(MOCK_JSON)

        assertEquals(3, flights.size, "expected 3 flights")

        // Sorted by price ascending: 159, 189, 245
        assertEquals(159, flights[0].price, "flights[0].price")
        assertEquals(189, flights[1].price, "flights[1].price")
        assertEquals(245, flights[2].price, "flights[2].price")

        // Cheapest is JetBlue B6 816
        assertEquals("JetBlue", flights[0].airline, "flights[0].airline")
        assertEquals("B6 816",  flights[0].flightNumber, "flights[0].flightNumber")
        assertEquals("5h 10m",  flights[0].duration, "flights[0].duration")
    }

    @Test
    fun testMissingSerpAPIKeyReturnsError() {
        Searcher.apiKeyOverride = "" // simulate missing key

        val ex = assertFailsWith<Exception> {
            Searcher.searchFlights("SFO", "JFK", "2024-01-15")
        }

        assertNotNull(ex.message, "error message should not be null")
        assertFalse(ex.message!!.isEmpty(), "error message should not be empty")
        assertTrue(ex.message!!.contains("SERPAPI_KEY"),
            "error message should mention SERPAPI_KEY, got: ${ex.message}")
    }
}
