package com.lakamsani.farefinder;

import java.util.HashMap;
import java.util.Map;

public class Airport {
    private static final Map<String, String> AIRPORTS = new HashMap<>();

    static {
        AIRPORTS.put("new york,ny", "JFK");
        AIRPORTS.put("los angeles,ca", "LAX");
        AIRPORTS.put("chicago,il", "ORD");
        AIRPORTS.put("houston,tx", "IAH");
        AIRPORTS.put("phoenix,az", "PHX");
        AIRPORTS.put("philadelphia,pa", "PHL");
        AIRPORTS.put("san antonio,tx", "SAT");
        AIRPORTS.put("san diego,ca", "SAN");
        AIRPORTS.put("dallas,tx", "DFW");
        AIRPORTS.put("san jose,ca", "SJC");
        AIRPORTS.put("austin,tx", "AUS");
        AIRPORTS.put("jacksonville,fl", "JAX");
        AIRPORTS.put("san francisco,ca", "SFO");
        AIRPORTS.put("columbus,oh", "CMH");
        AIRPORTS.put("charlotte,nc", "CLT");
        AIRPORTS.put("indianapolis,in", "IND");
        AIRPORTS.put("seattle,wa", "SEA");
        AIRPORTS.put("denver,co", "DEN");
        AIRPORTS.put("nashville,tn", "BNA");
        AIRPORTS.put("oklahoma city,ok", "OKC");
        AIRPORTS.put("el paso,tx", "ELP");
        AIRPORTS.put("washington,dc", "DCA");
        AIRPORTS.put("las vegas,nv", "LAS");
        AIRPORTS.put("louisville,ky", "SDF");
        AIRPORTS.put("baltimore,md", "BWI");
        AIRPORTS.put("milwaukee,wi", "MKE");
        AIRPORTS.put("albuquerque,nm", "ABQ");
        AIRPORTS.put("tucson,az", "TUS");
        AIRPORTS.put("fresno,ca", "FAT");
        AIRPORTS.put("sacramento,ca", "SMF");
        AIRPORTS.put("kansas city,mo", "MCI");
        AIRPORTS.put("atlanta,ga", "ATL");
        AIRPORTS.put("miami,fl", "MIA");
        AIRPORTS.put("minneapolis,mn", "MSP");
        AIRPORTS.put("portland,or", "PDX");
        AIRPORTS.put("detroit,mi", "DTW");
        AIRPORTS.put("boston,ma", "BOS");
        AIRPORTS.put("memphis,tn", "MEM");
        AIRPORTS.put("new orleans,la", "MSY");
        AIRPORTS.put("cleveland,oh", "CLE");
        AIRPORTS.put("tampa,fl", "TPA");
        AIRPORTS.put("orlando,fl", "MCO");
    }

    public static String lookupAirport(String city, String state) throws Exception {
        String key = city.strip().toLowerCase() + "," + state.strip().toLowerCase();
        String code = AIRPORTS.get(key);
        if (code == null) {
            throw new Exception("no airport found for " + city + ", " + state);
        }
        return code;
    }
}
