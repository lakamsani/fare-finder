"""CLI entry point for fare-finder (Python port)."""

import sys
from datetime import date, timedelta

from .airport import lookup_airport
from .searcher import search_flights


def _title_case(s: str) -> str:
    """Capitalise the first letter of each word."""
    return " ".join(word.capitalize() for word in s.split())


def main(argv: list[str] | None = None) -> None:
    """Parse CLI arguments and print the cheapest flight."""
    args = sys.argv[1:] if argv is None else argv

    if len(args) != 4:
        print(
            "Usage: fare-finder <city1> <state1> <city2> <state2>",
            file=sys.stderr,
        )
        print(
            'Example: fare-finder "San Francisco" CA "New York" NY',
            file=sys.stderr,
        )
        sys.exit(1)

    city1 = _title_case(args[0].strip())
    state1 = args[1].strip().upper()
    city2 = _title_case(args[2].strip())
    state2 = args[3].strip().upper()

    try:
        origin = lookup_airport(city1, state1)
    except ValueError as exc:
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)

    try:
        dest = lookup_airport(city2, state2)
    except ValueError as exc:
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)

    tomorrow = (date.today() + timedelta(days=1)).strftime("%Y-%m-%d")

    print(f"Searching for flights {origin} -> {dest} on {tomorrow}...\n")

    try:
        flights = search_flights(origin, dest, tomorrow)
    except Exception as exc:  # noqa: BLE001
        print(f"Error: {exc}", file=sys.stderr)
        sys.exit(1)

    if not flights:
        print(
            "No flights found for this route and date. "
            "Try a different date or city pair."
        )
        sys.exit(0)

    cheapest = flights[0]
    print(f"Cheapest flight: {origin} -> {dest}")
    print(f"{cheapest.airline} {cheapest.flight_number}")
    print(f"${cheapest.price}")
    print(f"Departs {cheapest.departure_time} | Arrives {cheapest.arrival_time}")
    print(f"Duration: {cheapest.duration}")


if __name__ == "__main__":
    main()
