package com.lakamsani.farefinder;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class AirportTest {

    @Test
    void testKnownAirports() throws Exception {
        Object[][] cases = {
            {"San Francisco", "CA", "SFO"},
            {"New York",      "NY", "JFK"},
            {"Chicago",       "IL", "ORD"},
            {"Atlanta",       "GA", "ATL"},
            {"Denver",        "CO", "DEN"},
            {"Seattle",       "WA", "SEA"},
            {"Miami",         "FL", "MIA"},
            {"Las Vegas",     "NV", "LAS"},
            {"Boston",        "MA", "BOS"},
            {"Dallas",        "TX", "DFW"},
        };

        for (Object[] tc : cases) {
            String city  = (String) tc[0];
            String state = (String) tc[1];
            String want  = (String) tc[2];

            String got = Airport.lookupAirport(city, state);
            assertEquals(want, got,
                    String.format("lookupAirport(%s, %s)", city, state));
        }
    }

    @Test
    void testUnknownCityReturnsError() {
        Exception ex = assertThrows(Exception.class,
                () -> Airport.lookupAirport("Nonexistent City", "XX"));
        assertNotNull(ex.getMessage());
    }
}
