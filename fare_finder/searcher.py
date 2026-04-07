import json
import os
import urllib.request
import urllib.parse

from .flight import Flight

BASE_URL = "https://serpapi.com/search"

# Overridable for testing
_api_key_override: str | None = None


def search_flights(origin: str, dest: str, date: str) -> list[Flight]:
    api_key = _api_key_override if _api_key_override is not None else os.environ.get("SERPAPI_KEY", "")

    if not api_key:
        raise RuntimeError(
            "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
        )

    params = urllib.parse.urlencode({
        "engine": "google_flights",
        "departure_id": origin,
        "arrival_id": dest,
        "outbound_date": date,
        "currency": "USD",
        "hl": "en",
        "api_key": api_key,
        "type": "2",
    })
    url = f"{BASE_URL}?{params}"

    with urllib.request.urlopen(url) as resp:
        if resp.status != 200:
            raise RuntimeError(f"API returned status {resp.status}")
        body = resp.read()

    return parse_flights(body)


def parse_flights(data: bytes) -> list[Flight]:
    root = json.loads(data)
    best = root.get("best_flights", [])
    other = root.get("other_flights", [])

    flights: list[Flight] = []
    for group in best + other:
        legs = group.get("flights", [])
        if not legs:
            continue
        first = legs[0]
        last = legs[-1]
        total_duration = sum(leg.get("duration", 0) for leg in legs)

        flights.append(Flight(
            airline=first.get("airline", ""),
            flight_number=first.get("flight_number", ""),
            price=group.get("price", 0),
            departure_time=first.get("departure_airport", {}).get("time", ""),
            arrival_time=last.get("arrival_airport", {}).get("time", ""),
            duration=_format_duration(total_duration),
        ))

    flights.sort(key=lambda f: f.price)
    return flights


def _format_duration(minutes: int) -> str:
    return f"{minutes // 60}h {minutes % 60}m"
