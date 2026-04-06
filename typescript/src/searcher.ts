import { Flight } from "./flight.js";

export let baseUrl = "https://serpapi.com/search";

/**
 * API key override.
 * undefined = read from SERPAPI_KEY environment variable
 * ""        = simulate missing key (triggers error)
 * other     = use this key directly
 *
 * Note: use the `state` object to mutate this from tests (ESM live bindings are read-only externally).
 */
export let apiKeyOverride: string | undefined = undefined;

/** Mutable state object — use this from tests to override baseUrl and apiKeyOverride. */
export const state: { baseUrl: string; apiKeyOverride: string | undefined } = {
  get baseUrl() { return baseUrl; },
  set baseUrl(v: string) { baseUrl = v; },
  get apiKeyOverride() { return apiKeyOverride; },
  set apiKeyOverride(v: string | undefined) { apiKeyOverride = v; },
};

export async function searchFlights(origin: string, dest: string, date: string): Promise<Flight[]> {
  let apiKey: string;
  if (apiKeyOverride !== undefined) {
    apiKey = apiKeyOverride;
  } else {
    apiKey = process.env["SERPAPI_KEY"] ?? "";
  }

  if (!apiKey) {
    throw new Error("SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com");
  }

  const url = `${baseUrl}?engine=google_flights&departure_id=${origin}&arrival_id=${dest}&outbound_date=${date}&currency=USD&hl=en&api_key=${apiKey}&type=2`;

  const response = await fetch(url);

  if (!response.ok) {
    throw new Error(`API returned status ${response.status}`);
  }

  const body = await response.text();
  return parseFlights(body);
}

export async function parseFlights(json: string): Promise<Flight[]> {
  let root: Record<string, unknown>;
  try {
    root = JSON.parse(json) as Record<string, unknown>;
  } catch (e) {
    throw new Error(`failed to parse JSON: ${(e as Error).message}`);
  }

  const flights: Flight[] = [];

  for (const groupKey of ["best_flights", "other_flights"]) {
    const groups = root[groupKey];
    if (!Array.isArray(groups)) continue;

    for (const group of groups as Record<string, unknown>[]) {
      const flightLegs = group["flights"];
      if (!Array.isArray(flightLegs) || flightLegs.length === 0) continue;

      const price = Number((group["price"] as number | undefined) ?? 0);

      const firstLeg = flightLegs[0] as Record<string, unknown>;
      const lastLeg = flightLegs[flightLegs.length - 1] as Record<string, unknown>;

      const airline = String((firstLeg["airline"] as string | undefined) ?? "");
      const flightNumber = String((firstLeg["flight_number"] as string | undefined) ?? "");
      const departureTime = String(
        ((firstLeg["departure_airport"] as Record<string, unknown> | undefined)?.["time"]) ?? ""
      );
      const arrivalTime = String(
        ((lastLeg["arrival_airport"] as Record<string, unknown> | undefined)?.["time"]) ?? ""
      );

      let totalDuration = 0;
      for (const leg of flightLegs as Record<string, unknown>[]) {
        totalDuration += Number((leg["duration"] as number | undefined) ?? 0);
      }

      flights.push(new Flight(airline, flightNumber, price, departureTime, arrivalTime, formatDuration(totalDuration)));
    }
  }

  flights.sort((a, b) => a.price - b.price);
  return flights;
}

function formatDuration(minutes: number): string {
  return `${Math.floor(minutes / 60)}h ${minutes % 60}m`;
}
