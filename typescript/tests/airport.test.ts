import { describe, it, expect } from "vitest";
import { lookupAirport } from "../src/airport.js";

describe("Airport", () => {
  it("testKnownAirports", () => {
    const cases: [string, string, string][] = [
      ["San Francisco", "CA", "SFO"],
      ["New York",      "NY", "JFK"],
      ["Chicago",       "IL", "ORD"],
      ["Atlanta",       "GA", "ATL"],
      ["Denver",        "CO", "DEN"],
      ["Seattle",       "WA", "SEA"],
      ["Miami",         "FL", "MIA"],
      ["Las Vegas",     "NV", "LAS"],
      ["Boston",        "MA", "BOS"],
      ["Dallas",        "TX", "DFW"],
    ];

    for (const [city, state, want] of cases) {
      expect(lookupAirport(city, state), `lookupAirport(${city}, ${state})`).toBe(want);
    }
  });

  it("testUnknownCityReturnsError", () => {
    expect(() => lookupAirport("Nonexistent City", "XX")).toThrow();
  });
});
