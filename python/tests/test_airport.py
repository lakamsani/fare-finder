"""Tests for the airport lookup module."""

import pytest

from fare_finder.airport import lookup_airport


@pytest.mark.parametrize(
    "city, state, expected",
    [
        ("San Francisco", "CA", "SFO"),
        ("New York",      "NY", "JFK"),
        ("Chicago",       "IL", "ORD"),
        ("Atlanta",       "GA", "ATL"),
        ("Denver",        "CO", "DEN"),
        ("Seattle",       "WA", "SEA"),
        ("Miami",         "FL", "MIA"),
        ("Las Vegas",     "NV", "LAS"),
        ("Boston",        "MA", "BOS"),
        ("Dallas",        "TX", "DFW"),
    ],
)
def test_known_airports(city: str, state: str, expected: str) -> None:
    assert lookup_airport(city, state) == expected


def test_unknown_city_raises_value_error() -> None:
    with pytest.raises(ValueError) as exc_info:
        lookup_airport("Nonexistent City", "XX")
    assert exc_info.value.args[0]  # message should be non-empty
