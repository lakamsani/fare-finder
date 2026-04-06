package com.lakamsani.farefinder

/**
 * Provides IATA airport code lookup by city and state.
 */
object Airport {

    private val airports: Map<String, String> = mapOf(
        "new york,ny"       to "JFK",
        "los angeles,ca"    to "LAX",
        "chicago,il"        to "ORD",
        "houston,tx"        to "IAH",
        "phoenix,az"        to "PHX",
        "philadelphia,pa"   to "PHL",
        "san antonio,tx"    to "SAT",
        "san diego,ca"      to "SAN",
        "dallas,tx"         to "DFW",
        "san jose,ca"       to "SJC",
        "austin,tx"         to "AUS",
        "jacksonville,fl"   to "JAX",
        "san francisco,ca"  to "SFO",
        "columbus,oh"       to "CMH",
        "charlotte,nc"      to "CLT",
        "indianapolis,in"   to "IND",
        "seattle,wa"        to "SEA",
        "denver,co"         to "DEN",
        "nashville,tn"      to "BNA",
        "oklahoma city,ok"  to "OKC",
        "el paso,tx"        to "ELP",
        "washington,dc"     to "DCA",
        "las vegas,nv"      to "LAS",
        "louisville,ky"     to "SDF",
        "baltimore,md"      to "BWI",
        "milwaukee,wi"      to "MKE",
        "albuquerque,nm"    to "ABQ",
        "tucson,az"         to "TUS",
        "fresno,ca"         to "FAT",
        "sacramento,ca"     to "SMF",
        "kansas city,mo"    to "MCI",
        "atlanta,ga"        to "ATL",
        "miami,fl"          to "MIA",
        "minneapolis,mn"    to "MSP",
        "portland,or"       to "PDX",
        "detroit,mi"        to "DTW",
        "boston,ma"         to "BOS",
        "memphis,tn"        to "MEM",
        "new orleans,la"    to "MSY",
        "cleveland,oh"      to "CLE",
        "tampa,fl"          to "TPA",
        "orlando,fl"        to "MCO",
    )

    /**
     * Returns the IATA code for a given city and state.
     *
     * @param city  the city name (case-insensitive)
     * @param state the two-letter state code (case-insensitive)
     * @return the IATA airport code
     * @throws Exception if no airport is found for the city/state combination
     */
    fun lookupAirport(city: String, state: String): String {
        val key = "${city.trim().lowercase()},${state.trim().lowercase()}"
        return airports[key] ?: throw Exception("no airport found for $city, $state")
    }
}
