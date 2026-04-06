import json
import os
import urllib.request
from typing import Union

from fare_finder.flight import Flight

BASE_URL = "https://serpapi.com/search"

_api_key_override: Union[str, None] = None


def search_flights(origin: str, dest: str, date: str) -> list[Flight]:
    if _api_key_override is not None:
        api_key = _api_key_override
    else:
        api_key = os.environ.get("SERPAPI_KEY", "")

    if not api_key:
        raise RuntimeError(
            "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
        )

    url = (
        f"{BASE_URL}?engine=google_flights&departure_id={origin}&arrival_id={dest}"
        f"&outbound_date={date}&currency=USD&hl=en&api_key={api_key}&type=2"
    )

    with urllib.request.urlopen(url) as resp:
        if resp.status != 200:
            raise RuntimeError(f"API returned status {resp.status}")
        body = resp.read()

    return parse_flights(body)


def parse_flights(data: Union[str, bytes]) -> list[Flight]:
    root = json.loads(data)

    best = root.get("best_flights", [])
    other = root.get("other_flights", [])

    flights = []
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
