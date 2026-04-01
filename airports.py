"""Airport lookup: maps city/state to IATA airport codes."""

AIRPORT_MAP = {
    "new york,ny": "JFK",
    "los angeles,ca": "LAX",
    "chicago,il": "ORD",
    "houston,tx": "IAH",
    "phoenix,az": "PHX",
    "philadelphia,pa": "PHL",
    "san antonio,tx": "SAT",
    "san diego,ca": "SAN",
    "dallas,tx": "DFW",
    "san jose,ca": "SJC",
    "austin,tx": "AUS",
    "jacksonville,fl": "JAX",
    "san francisco,ca": "SFO",
    "columbus,oh": "CMH",
    "charlotte,nc": "CLT",
    "indianapolis,in": "IND",
    "seattle,wa": "SEA",
    "denver,co": "DEN",
    "nashville,tn": "BNA",
    "oklahoma city,ok": "OKC",
    "el paso,tx": "ELP",
    "washington,dc": "DCA",
    "las vegas,nv": "LAS",
    "louisville,ky": "SDF",
    "baltimore,md": "BWI",
    "milwaukee,wi": "MKE",
    "albuquerque,nm": "ABQ",
    "tucson,az": "TUS",
    "fresno,ca": "FAT",
    "sacramento,ca": "SMF",
    "kansas city,mo": "MCI",
    "atlanta,ga": "ATL",
    "miami,fl": "MIA",
    "minneapolis,mn": "MSP",
    "portland,or": "PDX",
    "detroit,mi": "DTW",
    "boston,ma": "BOS",
    "memphis,tn": "MEM",
    "new orleans,la": "MSY",
    "cleveland,oh": "CLE",
    "tampa,fl": "TPA",
    "orlando,fl": "MCO",
}


def lookup_airport(city: str, state: str) -> str:
    """Return the IATA code for a city and state.

    Raises ValueError if the city/state pair is not found.
    """
    key = f"{city.strip().lower()},{state.strip().lower()}"
    code = AIRPORT_MAP.get(key)
    if code is None:
        raise ValueError(f"no airport found for {city}, {state}")
    return code
