"""Tests for airport lookup."""

import pytest

import sys
import os
sys.path.insert(0, os.path.dirname(os.path.dirname(__file__)))

from airports import lookup_airport


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
def test_airport_lookup(city, state, expected):
    assert lookup_airport(city, state) == expected


def test_airport_lookup_not_found():
    with pytest.raises(ValueError, match="no airport found"):
        lookup_airport("Nonexistent City", "XX")
