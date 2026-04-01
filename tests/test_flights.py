"""Tests for flight search and parsing."""

import os
import sys
from unittest.mock import patch, MagicMock

import pytest

sys.path.insert(0, os.path.dirname(os.path.dirname(__file__)))

import flights as flights_module
from flights import Flight, parse_flights, search_flights


MOCK_RESPONSE = {
    "best_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 08:05"},
                    "arrival_airport": {"time": "2024-01-15 16:23"},
                    "duration": 318,
                    "airline": "United Airlines",
                    "flight_number": "UA 101",
                }
            ],
            "price": 189,
        }
    ],
    "other_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 10:30"},
                    "arrival_airport": {"time": "2024-01-15 18:45"},
                    "duration": 315,
                    "airline": "Delta Air Lines",
                    "flight_number": "DL 405",
                }
            ],
            "price": 245,
        },
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 06:00"},
                    "arrival_airport": {"time": "2024-01-15 14:10"},
                    "duration": 310,
                    "airline": "JetBlue",
                    "flight_number": "B6 816",
                }
            ],
            "price": 159,
        },
    ],
}


def test_flight_sorting():
    unsorted = [
        Flight("Delta", "DL 1", 350, "", "", ""),
        Flight("Spirit", "NK 1", 89, "", "", ""),
        Flight("United", "UA 1", 210, "", "", ""),
        Flight("JetBlue", "B6 1", 175, "", "", ""),
    ]
    unsorted.sort(key=lambda f: f.price)

    assert unsorted[0].price == 89
    assert unsorted[0].airline == "Spirit"
    for i in range(1, len(unsorted)):
        assert unsorted[i].price >= unsorted[i - 1].price


def test_search_flights_no_key():
    os.environ.pop("SERPAPI_KEY", None)
    with pytest.raises(ValueError, match="SERPAPI_KEY"):
        search_flights("SFO", "JFK", "2024-01-15")


def test_flight_parsing():
    result = parse_flights(MOCK_RESPONSE)

    assert len(result) == 3

    # Sorted by price: 159, 189, 245
    assert result[0].price == 159
    assert result[0].airline == "JetBlue"
    assert result[0].flight_number == "B6 816"
    assert result[0].duration == "5h 10m"

    assert result[1].price == 189
    assert result[2].price == 245


def test_search_flights_with_mock_server():
    mock_resp = MagicMock()
    mock_resp.json.return_value = MOCK_RESPONSE
    mock_resp.raise_for_status = MagicMock()

    os.environ["SERPAPI_KEY"] = "test-key"
    try:
        with patch("flights.requests.get", return_value=mock_resp):
            result = search_flights("SFO", "JFK", "2024-01-15")

        assert len(result) == 3
        assert result[0].price == 159
        assert result[1].price == 189
        assert result[2].price == 245
    finally:
        os.environ.pop("SERPAPI_KEY", None)
