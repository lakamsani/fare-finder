package com.lakamsani.farefinder;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class AirportTest {

    @Test
    public void testKnownAirports() throws Exception {
        assertEquals("SFO", Airport.lookupAirport("San Francisco", "CA"));
        assertEquals("JFK", Airport.lookupAirport("New York", "NY"));
        assertEquals("ORD", Airport.lookupAirport("Chicago", "IL"));
        assertEquals("ATL", Airport.lookupAirport("Atlanta", "GA"));
        assertEquals("DEN", Airport.lookupAirport("Denver", "CO"));
        assertEquals("SEA", Airport.lookupAirport("Seattle", "WA"));
        assertEquals("MIA", Airport.lookupAirport("Miami", "FL"));
        assertEquals("LAS", Airport.lookupAirport("Las Vegas", "NV"));
        assertEquals("BOS", Airport.lookupAirport("Boston", "MA"));
        assertEquals("DFW", Airport.lookupAirport("Dallas", "TX"));
    }

    @Test
    public void testUnknownCityReturnsError() {
        assertThrows(Exception.class, () -> Airport.lookupAirport("Nonexistent City", "XX"));
    }
}
