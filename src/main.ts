import { lookupAirport } from "./airport.js";
import { searchFlights } from "./searcher.js";

function titleCase(s: string): string {
  return s
    .split(/\s+/)
    .map((w) => (w.length > 0 ? w[0].toUpperCase() + w.slice(1).toLowerCase() : w))
    .join(" ");
}

async function main(): Promise<void> {
  const args = process.argv.slice(2);

  if (args.length !== 4) {
    console.error("Usage: fare-finder <city1> <state1> <city2> <state2>");
    console.error('Example: fare-finder "San Francisco" CA "New York" NY');
    process.exit(1);
  }

  const city1 = titleCase(args[0].trim());
  const state1 = args[1].trim().toUpperCase();
  const city2 = titleCase(args[2].trim());
  const state2 = args[3].trim().toUpperCase();

  const origin = lookupAirport(city1, state1);
  const dest = lookupAirport(city2, state2);

  const tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);
  const date = tomorrow.toISOString().split("T")[0];

  console.log(`Searching for flights ${origin} -> ${dest} on ${date}...\n`);

  const flights = await searchFlights(origin, dest, date);

  if (flights.length === 0) {
    console.log(
      "No flights found for this route and date. Try a different date or city pair."
    );
    process.exit(0);
  }

  const cheapest = flights[0];
  console.log(`Cheapest flight: ${origin} -> ${dest}`);
  console.log(`${cheapest.airline} ${cheapest.flightNumber}`);
  console.log(`$${cheapest.price}`);
  console.log(`Departs ${cheapest.departureTime} | Arrives ${cheapest.arrivalTime}`);
  console.log(`Duration: ${cheapest.duration}`);
}

main().catch((err) => {
  console.error(`Error: ${err.message}`);
  process.exit(1);
});
