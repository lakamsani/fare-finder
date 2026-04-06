"""Handles flight search via SerpAPI and JSON parsing."""

import json
import os
import urllib.request
import urllib.parse
from typing import Optional

from .flight import Flight

# SerpAPI endpoint — can be overridden in tests.
_BASE_URL = "https://serpapi.com/search"

# API key override for testing.
# None  = read from SERPAPI_KEY env var
# ""    = simulate missing key (triggers error)
# other = use directly
_api_key_override: Optional[str] = None


def _set_base_url(url: str) -> None:
    """Override the SerpAPI base URL (for tests)."""
    global _BASE_URL
    _BASE_URL = url


def _set_api_key_override(key: Optional[str]) -> None:
    """Override the API key (for tests). Pass None to reset."""
    global _api_key_override
    _api_key_override = key


def search_flights(origin: str, dest: str, date: str) -> list[Flight]:
    """Fetch flights from SerpAPI for the given origin, destination, and date.

    Args:
        origin: 3-letter IATA origin code.
        dest:   3-letter IATA destination code.
        date:   Departure date in ``YYYY-MM-DD`` format.

    Returns:
        List of :class:`Flight` objects sorted by price ascending.

    Raises:
        ValueError: If the SerpAPI key is missing.
        RuntimeError: If the HTTP request fails or returns a non-200 status.
    """
    if _api_key_override is not None:
        api_key = _api_key_override
    else:
        api_key = os.environ.get("SERPAPI_KEY", "")

    if not api_key:
        raise ValueError(
            "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
        )

    params = urllib.parse.urlencode(
        {
            "engine": "google_flights",
            "departure_id": origin,
            "arrival_id": dest,
            "outbound_date": date,
            "currency": "USD",
            "hl": "en",
            "api_key": api_key,
            "type": "2",
        }
    )
    url = f"{_BASE_URL}?{params}"

    req = urllib.request.Request(url)
    try:
        with urllib.request.urlopen(req) as resp:
            if resp.status != 200:
                raise RuntimeError(f"API returned status {resp.status}")
            body = resp.read().decode("utf-8")
    except urllib.error.HTTPError as exc:
        raise RuntimeError(f"API returned status {exc.code}") from exc
    except urllib.error.URLError as exc:
        raise RuntimeError(f"HTTP request failed: {exc.reason}") from exc

    return parse_flights(body)


def parse_flights(json_text: str) -> list[Flight]:
    """Parse a SerpAPI JSON response and return flights sorted by price ascending.

    Args:
        json_text: Raw JSON string from SerpAPI.

    Returns:
        List of :class:`Flight` objects sorted by price ascending.

    Raises:
        ValueError: If JSON parsing fails.
    """
    try:
        data = json.loads(json_text)
    except json.JSONDecodeError as exc:
        raise ValueError(f"failed to parse JSON: {exc}") from exc

    flights: list[Flight] = []

    for group_key in ("best_flights", "other_flights"):
        groups = data.get(group_key, [])
        if not isinstance(groups, list):
            continue
        for group in groups:
            legs = group.get("flights", [])
            if not legs:
                continue
            price = int(group.get("price", 0))

            first_leg = legs[0]
            last_leg = legs[-1]

            airline = first_leg.get("airline", "")
            flight_number = first_leg.get("flight_number", "")
            departure_time = first_leg.get("departure_airport", {}).get("time", "")
            arrival_time = last_leg.get("arrival_airport", {}).get("time", "")

            total_duration = sum(leg.get("duration", 0) for leg in legs)

            flights.append(
                Flight(
                    airline=airline,
                    flight_number=flight_number,
                    price=price,
                    departure_time=departure_time,
                    arrival_time=arrival_time,
                    duration=_format_duration(total_duration),
                )
            )

    flights.sort(key=lambda f: f.price)
    return flights


def _format_duration(minutes: int) -> str:
    """Convert total minutes into a human-readable ``Xh Ym`` string."""
    return f"{minutes // 60}h {minutes % 60}m"
