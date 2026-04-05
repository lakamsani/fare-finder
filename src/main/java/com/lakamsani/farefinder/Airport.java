package com.lakamsani.farefinder;

import java.util.HashMap;
import java.util.Map;

/**
 * Provides IATA airport code lookup by city and state.
 */
public class Airport {

    private static final Map<String, String> airports = new HashMap<>();

    static {
        airports.put("new york,ny",       "JFK");
        airports.put("los angeles,ca",    "LAX");
        airports.put("chicago,il",        "ORD");
        airports.put("houston,tx",        "IAH");
        airports.put("phoenix,az",        "PHX");
        airports.put("philadelphia,pa",   "PHL");
        airports.put("san antonio,tx",    "SAT");
        airports.put("san diego,ca",      "SAN");
        airports.put("dallas,tx",         "DFW");
        airports.put("san jose,ca",       "SJC");
        airports.put("austin,tx",         "AUS");
        airports.put("jacksonville,fl",   "JAX");
        airports.put("san francisco,ca",  "SFO");
        airports.put("columbus,oh",       "CMH");
        airports.put("charlotte,nc",      "CLT");
        airports.put("indianapolis,in",   "IND");
        airports.put("seattle,wa",        "SEA");
        airports.put("denver,co",         "DEN");
        airports.put("nashville,tn",      "BNA");
        airports.put("oklahoma city,ok",  "OKC");
        airports.put("el paso,tx",        "ELP");
        airports.put("washington,dc",     "DCA");
        airports.put("las vegas,nv",      "LAS");
        airports.put("louisville,ky",     "SDF");
        airports.put("baltimore,md",      "BWI");
        airports.put("milwaukee,wi",      "MKE");
        airports.put("albuquerque,nm",    "ABQ");
        airports.put("tucson,az",         "TUS");
        airports.put("fresno,ca",         "FAT");
        airports.put("sacramento,ca",     "SMF");
        airports.put("kansas city,mo",    "MCI");
        airports.put("atlanta,ga",        "ATL");
        airports.put("miami,fl",          "MIA");
        airports.put("minneapolis,mn",    "MSP");
        airports.put("portland,or",       "PDX");
        airports.put("detroit,mi",        "DTW");
        airports.put("boston,ma",         "BOS");
        airports.put("memphis,tn",        "MEM");
        airports.put("new orleans,la",    "MSY");
        airports.put("cleveland,oh",      "CLE");
        airports.put("tampa,fl",          "TPA");
        airports.put("orlando,fl",        "MCO");
    }

    /**
     * Returns the IATA code for a given city and state.
     *
     * @param city  the city name (case-insensitive)
     * @param state the two-letter state code (case-insensitive)
     * @return the IATA airport code
     * @throws Exception if no airport is found for the city/state combination
     */
    public static String lookupAirport(String city, String state) throws Exception {
        String key = city.trim().toLowerCase() + "," + state.trim().toLowerCase();
        String code = airports.get(key);
        if (code == null) {
            throw new Exception("no airport found for " + city + ", " + state);
        }
        return code;
    }
}
