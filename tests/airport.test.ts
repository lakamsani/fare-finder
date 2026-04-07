import { describe, it, expect } from "vitest";
import { lookupAirport } from "../src/airport.js";

describe("lookupAirport", () => {
  const knownCities: [string, string, string][] = [
    ["San Francisco", "CA", "SFO"],
    ["New York", "NY", "JFK"],
    ["Chicago", "IL", "ORD"],
    ["Atlanta", "GA", "ATL"],
    ["Denver", "CO", "DEN"],
    ["Seattle", "WA", "SEA"],
    ["Miami", "FL", "MIA"],
    ["Las Vegas", "NV", "LAS"],
    ["Boston", "MA", "BOS"],
    ["Dallas", "TX", "DFW"],
  ];

  it.each(knownCities)(
    "returns correct code for %s, %s",
    (city, state, expected) => {
      expect(lookupAirport(city, state)).toBe(expected);
    }
  );

  it("throws for unknown city", () => {
    expect(() => lookupAirport("Nonexistent City", "XX")).toThrow(
      "No airport found for Nonexistent City, XX"
    );
  });
});
