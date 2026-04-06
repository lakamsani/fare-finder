package com.lakamsani.farefinder

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertFailsWith
import kotlin.test.assertNotNull

class AirportTest {

    @Test
    fun testKnownAirports() {
        val cases = listOf(
            Triple("San Francisco", "CA", "SFO"),
            Triple("New York",      "NY", "JFK"),
            Triple("Chicago",       "IL", "ORD"),
            Triple("Atlanta",       "GA", "ATL"),
            Triple("Denver",        "CO", "DEN"),
            Triple("Seattle",       "WA", "SEA"),
            Triple("Miami",         "FL", "MIA"),
            Triple("Las Vegas",     "NV", "LAS"),
            Triple("Boston",        "MA", "BOS"),
            Triple("Dallas",        "TX", "DFW"),
        )

        for ((city, state, want) in cases) {
            val got = Airport.lookupAirport(city, state)
            assertEquals(want, got, "lookupAirport($city, $state)")
        }
    }

    @Test
    fun testUnknownCityReturnsError() {
        val ex = assertFailsWith<Exception> {
            Airport.lookupAirport("Nonexistent City", "XX")
        }
        assertNotNull(ex.message)
    }
}
