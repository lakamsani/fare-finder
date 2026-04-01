"""Flight search and parsing using SerpAPI Google Flights."""

import os
from dataclasses import dataclass
from typing import Any

import requests

SERP_API_BASE_URL = "https://serpapi.com/search"


@dataclass
class Flight:
    airline: str
    flight_number: str
    price: int
    departure_time: str
    arrival_time: str
    duration: str


def search_flights(origin: str, dest: str, date: str) -> list[Flight]:
    """Query SerpAPI for flights and return them sorted by price (cheapest first)."""
    api_key = os.environ.get("SERPAPI_KEY")
    if not api_key:
        raise ValueError(
            "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
        )

    params = {
        "engine": "google_flights",
        "departure_id": origin,
        "arrival_id": dest,
        "outbound_date": date,
        "currency": "USD",
        "hl": "en",
        "api_key": api_key,
        "type": "2",
    }

    resp = requests.get(SERP_API_BASE_URL, params=params, timeout=30)
    resp.raise_for_status()

    return parse_flights(resp.json())


def parse_flights(data: dict[str, Any]) -> list[Flight]:
    """Parse the SerpAPI JSON response into a sorted list of flights."""
    flights: list[Flight] = []

    all_groups = data.get("best_flights", []) + data.get("other_flights", [])
    for group in all_groups:
        legs = group.get("flights", [])
        if not legs:
            continue
        first_leg = legs[0]
        total_minutes = sum(leg.get("duration", 0) for leg in legs)
        flights.append(
            Flight(
                airline=first_leg.get("airline", ""),
                flight_number=first_leg.get("flight_number", ""),
                price=group.get("price", 0),
                departure_time=first_leg.get("departure_airport", {}).get("time", ""),
                arrival_time=legs[-1].get("arrival_airport", {}).get("time", ""),
                duration=_format_duration(total_minutes),
            )
        )

    flights.sort(key=lambda f: f.price)
    return flights


def _format_duration(minutes: int) -> str:
    return f"{minutes // 60}h {minutes % 60}m"
