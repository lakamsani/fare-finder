package com.fareFinder;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class AirportTest {

    @Test
    void testKnownAirports() {
        assertEquals("SFO", Airport.lookup("San Francisco", "CA"));
        assertEquals("JFK", Airport.lookup("New York", "NY"));
        assertEquals("ORD", Airport.lookup("Chicago", "IL"));
        assertEquals("ATL", Airport.lookup("Atlanta", "GA"));
        assertEquals("DEN", Airport.lookup("Denver", "CO"));
        assertEquals("SEA", Airport.lookup("Seattle", "WA"));
        assertEquals("MIA", Airport.lookup("Miami", "FL"));
        assertEquals("LAS", Airport.lookup("Las Vegas", "NV"));
        assertEquals("BOS", Airport.lookup("Boston", "MA"));
        assertEquals("DFW", Airport.lookup("Dallas", "TX"));
    }

    @Test
    void testUnknownCityThrows() {
        assertThrows(IllegalArgumentException.class,
            () -> Airport.lookup("Nonexistent City", "XX"));
    }
}
