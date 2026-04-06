import pytest

import fare_finder.searcher as searcher
from fare_finder.searcher import parse_flights, search_flights

MOCK_JSON = """
{
    "best_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 08:05"},
                    "arrival_airport": {"time": "2024-01-15 16:23"},
                    "duration": 318,
                    "airline": "United Airlines",
                    "flight_number": "UA 101"
                }
            ],
            "price": 189
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
                    "flight_number": "DL 405"
                }
            ],
            "price": 245
        },
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 06:00"},
                    "arrival_airport": {"time": "2024-01-15 14:10"},
                    "duration": 310,
                    "airline": "JetBlue",
                    "flight_number": "B6 816"
                }
            ],
            "price": 159
        }
    ]
}
"""


def test_parse_flights():
    flights = parse_flights(MOCK_JSON)

    assert len(flights) == 3

    # Sorted by price ascending: 159, 189, 245
    assert flights[0].price == 159
    assert flights[1].price == 189
    assert flights[2].price == 245

    # Cheapest is JetBlue B6 816
    assert flights[0].airline == "JetBlue"
    assert flights[0].flight_number == "B6 816"
    assert flights[0].duration == "5h 10m"


def test_missing_serpapi_key_raises():
    searcher._api_key_override = ""
    try:
        with pytest.raises(Exception, match="SERPAPI_KEY"):
            search_flights("SFO", "JFK", "2024-01-15")
    finally:
        searcher._api_key_override = None
