import { lookupAirport } from "./airport.js";
import { searchFlights } from "./searcher.js";

function titleCase(s: string): string {
  if (!s) return s;
  let capitalizeNext = true;
  let result = "";
  for (const c of s) {
    if (/\s/.test(c)) {
      capitalizeNext = true;
      result += c;
    } else if (capitalizeNext) {
      result += c.toUpperCase();
      capitalizeNext = false;
    } else {
      result += c.toLowerCase();
    }
  }
  return result;
}

async function main(): Promise<void> {
  const args = process.argv.slice(2);
  if (args.length !== 4) {
    process.stderr.write("Usage: fare-finder <city1> <state1> <city2> <state2>\n");
    process.stderr.write('Example: fare-finder "San Francisco" CA "New York" NY\n');
    process.exit(1);
  }

  const city1 = titleCase(args[0].trim());
  const state1 = args[1].trim().toUpperCase();
  const city2 = titleCase(args[2].trim());
  const state2 = args[3].trim().toUpperCase();

  let origin: string;
  try {
    origin = lookupAirport(city1, state1);
  } catch (e) {
    process.stderr.write(`Error: ${(e as Error).message}\n`);
    process.exit(1);
  }

  let dest: string;
  try {
    dest = lookupAirport(city2, state2);
  } catch (e) {
    process.stderr.write(`Error: ${(e as Error).message}\n`);
    process.exit(1);
  }

  const tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);
  const date = tomorrow.toISOString().split("T")[0];

  process.stdout.write(`Searching for flights ${origin} -> ${dest} on ${date}...\n\n`);

  let flights;
  try {
    flights = await searchFlights(origin, dest, date);
  } catch (e) {
    process.stderr.write(`Error: ${(e as Error).message}\n`);
    process.exit(1);
  }

  if (flights.length === 0) {
    process.stdout.write("No flights found for this route and date. Try a different date or city pair.\n");
    process.exit(0);
  }

  const cheapest = flights[0];
  process.stdout.write(`Cheapest flight: ${origin} -> ${dest}\n`);
  process.stdout.write(`${cheapest.airline} ${cheapest.flightNumber}\n`);
  process.stdout.write(`$${cheapest.price}\n`);
  process.stdout.write(`Departs ${cheapest.departureTime} | Arrives ${cheapest.arrivalTime}\n`);
  process.stdout.write(`Duration: ${cheapest.duration}\n`);
}

main();
