import { describe, it, expect, beforeAll, afterAll, vi } from "vitest";
import http from "node:http";
import { parseFlights, searchFlights, setBaseURL } from "../src/searcher.js";

const mockJSON = `{
  "best_flights": [
    {
      "flights": [
        {
          "departure_airport": {"time": "2024-01-15 08:05"},
          "arrival_airport":   {"time": "2024-01-15 16:23"},
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
          "arrival_airport":   {"time": "2024-01-15 18:45"},
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
          "arrival_airport":   {"time": "2024-01-15 14:10"},
          "duration": 310,
          "airline": "JetBlue",
          "flight_number": "B6 816"
        }
      ],
      "price": 159
    }
  ]
}`;

describe("parseFlights", () => {
  it("parses valid JSON and sorts by price", () => {
    const flights = parseFlights(mockJSON);

    expect(flights).toHaveLength(3);
    expect(flights[0].price).toBe(159);
    expect(flights[1].price).toBe(189);
    expect(flights[2].price).toBe(245);

    expect(flights[0].airline).toBe("JetBlue");
    expect(flights[0].flightNumber).toBe("B6 816");
    expect(flights[0].duration).toBe("5h 10m");
  });

  it("returns empty array for empty JSON", () => {
    const flights = parseFlights("{}");
    expect(flights).toHaveLength(0);
  });

  it("throws for invalid JSON", () => {
    expect(() => parseFlights("not valid json")).toThrow();
  });
});

describe("searchFlights", () => {
  it("throws when SERPAPI_KEY is missing", async () => {
    const original = process.env.SERPAPI_KEY;
    delete process.env.SERPAPI_KEY;

    await expect(searchFlights("SFO", "JFK", "2024-01-15")).rejects.toThrow(
      "SERPAPI_KEY"
    );

    if (original !== undefined) process.env.SERPAPI_KEY = original;
  });

  it("fetches and parses flights from HTTP server", async () => {
    const server = http.createServer((_req, res) => {
      res.writeHead(200, { "Content-Type": "application/json" });
      res.end(mockJSON);
    });

    await new Promise<void>((resolve) => server.listen(0, resolve));
    const addr = server.address() as { port: number };

    const origKey = process.env.SERPAPI_KEY;
    process.env.SERPAPI_KEY = "test-key";
    setBaseURL(`http://localhost:${addr.port}`);

    try {
      const flights = await searchFlights("SFO", "JFK", "2024-01-15");
      expect(flights).toHaveLength(3);
    } finally {
      setBaseURL("https://serpapi.com/search");
      if (origKey !== undefined) {
        process.env.SERPAPI_KEY = origKey;
      } else {
        delete process.env.SERPAPI_KEY;
      }
      server.close();
    }
  });
});
