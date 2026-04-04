package com.fareFinder;

import org.junit.Test;
import static org.junit.Assert.*;

public class AirportTest {

    @Test
    public void testKnownAirports() {
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

    @Test(expected = IllegalArgumentException.class)
    public void testUnknownCityThrows() {
        Airport.lookup("Nonexistent City", "XX");
    }
}
