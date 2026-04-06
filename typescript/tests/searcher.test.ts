import { describe, it, expect, afterEach } from "vitest";
import { parseFlights, searchFlights, state } from "../src/searcher.js";

const MOCK_JSON = `{
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
}`;

afterEach(() => {
  state.apiKeyOverride = undefined;
});

describe("Searcher", () => {
  it("testParseFlights", async () => {
    const flights = await parseFlights(MOCK_JSON);

    expect(flights.length).toBe(3);

    // Sorted by price ascending: 159, 189, 245
    expect(flights[0].price).toBe(159);
    expect(flights[1].price).toBe(189);
    expect(flights[2].price).toBe(245);

    // Cheapest is JetBlue B6 816
    expect(flights[0].airline).toBe("JetBlue");
    expect(flights[0].flightNumber).toBe("B6 816");
    expect(flights[0].duration).toBe("5h 10m");
  });

  it("testMissingSerpAPIKeyReturnsError", async () => {
    state.apiKeyOverride = ""; // simulate missing key

    await expect(searchFlights("SFO", "JFK", "2024-01-15")).rejects.toThrow("SERPAPI_KEY");
  });
});
