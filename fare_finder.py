#!/usr/bin/env python3
"""fare-finder: find the cheapest US domestic flight between two cities."""

import argparse
import sys
from datetime import date, timedelta

from airports import lookup_airport
from flights import search_flights


def title_case(s: str) -> str:
    return s.strip().title()


def main() -> None:
    parser = argparse.ArgumentParser(
        prog="fare-finder",
        description="Find the cheapest flight between two US cities.",
    )
    parser.add_argument("city1", help="Origin city (e.g. 'San Francisco')")
    parser.add_argument("state1", help="Origin state abbreviation (e.g. CA)")
    parser.add_argument("city2", help="Destination city (e.g. 'New York')")
    parser.add_argument("state2", help="Destination state abbreviation (e.g. NY)")
    args = parser.parse_args()

    city1 = title_case(args.city1)
    state1 = args.state1.strip().upper()
    city2 = title_case(args.city2)
    state2 = args.state2.strip().upper()

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

    tomorrow = (date.today() + timedelta(days=1)).strftime("%Y-%m-%d")

    print(f"Searching for flights {origin} \u2192 {dest} on {tomorrow}...\n")

    try:
        flights = search_flights(origin, dest, tomorrow)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    if not flights:
        print("No flights found for this route and date. Try a different date or city pair.")
        sys.exit(0)

    cheapest = flights[0]
    print(f"\U0001f6eb Cheapest flight: {origin} \u2192 {dest}")
    print(f"\u2708\ufe0f  {cheapest.airline} {cheapest.flight_number}")
    print(f"\U0001f4b0 ${cheapest.price}")
    print(f"\U0001f557 Departs {cheapest.departure_time} | Arrives {cheapest.arrival_time}")
    print(f"\u23f1\ufe0f  Duration: {cheapest.duration}")


if __name__ == "__main__":
    main()
