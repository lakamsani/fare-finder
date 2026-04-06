import pytest

from fare_finder.airport import lookup_airport


@pytest.mark.parametrize("city,state,expected", [
    ("San Francisco", "CA", "SFO"),
    ("New York", "NY", "JFK"),
    ("Chicago", "IL", "ORD"),
    ("Atlanta", "GA", "ATL"),
    ("Denver", "CO", "DEN"),
    ("Seattle", "WA", "SEA"),
    ("Miami", "FL", "MIA"),
    ("Las Vegas", "NV", "LAS"),
    ("Boston", "MA", "BOS"),
    ("Dallas", "TX", "DFW"),
])
def test_known_airports(city, state, expected):
    assert lookup_airport(city, state) == expected


def test_unknown_city_raises():
    with pytest.raises(ValueError):
        lookup_airport("Nonexistent City", "XX")
