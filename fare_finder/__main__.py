import sys
from datetime import date, timedelta

from .airport import lookup_airport
from .searcher import search_flights


def _title_case(s: str) -> str:
    return " ".join(word.capitalize() for word in s.split())


def main() -> None:
    if len(sys.argv) != 5:
        print("Usage: fare-finder <city1> <state1> <city2> <state2>", file=sys.stderr)
        print('Example: fare-finder "San Francisco" CA "New York" NY', file=sys.stderr)
        sys.exit(1)

    city1 = _title_case(sys.argv[1].strip())
    state1 = sys.argv[2].strip().upper()
    city2 = _title_case(sys.argv[3].strip())
    state2 = sys.argv[4].strip().upper()

    try:
        origin = lookup_airport(city1, state1)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    try:
        dest = lookup_airport(city2, state2)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    tomorrow = (date.today() + timedelta(days=1)).isoformat()

    print(f"Searching for flights {origin} -> {dest} on {tomorrow}...\n")

    try:
        flights = search_flights(origin, dest, tomorrow)
    except RuntimeError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    if not flights:
        print("No flights found for this route and date. Try a different date or city pair.")
        sys.exit(0)

    cheapest = flights[0]
    print(f"Cheapest flight: {origin} -> {dest}")
    print(f"{cheapest.airline} {cheapest.flight_number}")
    print(f"${cheapest.price}")
    print(f"Departs {cheapest.departure_time} | Arrives {cheapest.arrival_time}")
    print(f"Duration: {cheapest.duration}")


if __name__ == "__main__":
    main()
