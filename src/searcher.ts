import { Flight } from "./flight.js";

export let baseURL = "https://serpapi.com/search";

export function setBaseURL(url: string): void {
  baseURL = url;
}

interface AirportTime {
  time: string;
}

interface LegData {
  departure_airport: AirportTime;
  arrival_airport: AirportTime;
  duration: number;
  airline: string;
  flight_number: string;
}

interface FlightGroup {
  flights: LegData[];
  price: number;
}

interface SerpAPIResponse {
  best_flights?: FlightGroup[];
  other_flights?: FlightGroup[];
}

function formatDuration(minutes: number): string {
  return `${Math.floor(minutes / 60)}h ${minutes % 60}m`;
}

export function parseFlights(jsonText: string): Flight[] {
  const data: SerpAPIResponse = JSON.parse(jsonText);
  const flights: Flight[] = [];

  for (const groups of [data.best_flights ?? [], data.other_flights ?? []]) {
    for (const group of groups) {
      const legs = group.flights;
      if (legs.length === 0) continue;

      const firstLeg = legs[0];
      const lastLeg = legs[legs.length - 1];
      const totalDuration = legs.reduce((sum, leg) => sum + leg.duration, 0);

      flights.push({
        airline: firstLeg.airline,
        flightNumber: firstLeg.flight_number,
        price: group.price,
        departureTime: firstLeg.departure_airport.time,
        arrivalTime: lastLeg.arrival_airport.time,
        duration: formatDuration(totalDuration),
      });
    }
  }

  flights.sort((a, b) => a.price - b.price);
  return flights;
}

export async function searchFlights(
  origin: string,
  dest: string,
  date: string
): Promise<Flight[]> {
  const apiKey = process.env.SERPAPI_KEY;
  if (!apiKey) {
    throw new Error(
      "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
    );
  }

  const params = new URLSearchParams({
    engine: "google_flights",
    departure_id: origin,
    arrival_id: dest,
    outbound_date: date,
    currency: "USD",
    hl: "en",
    api_key: apiKey,
    type: "2",
  });

  const resp = await fetch(`${baseURL}?${params}`);
  if (!resp.ok) {
    throw new Error(`API returned status ${resp.status}`);
  }

  const body = await resp.text();
  return parseFlights(body);
}
