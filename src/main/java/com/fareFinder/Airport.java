package com.fareFinder;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

public class Airport {
    private static final Map<String, String> AIRPORTS;

    static {
        Map<String, String> map = new HashMap<>();
        map.put("new york,ny", "JFK");
        map.put("los angeles,ca", "LAX");
        map.put("chicago,il", "ORD");
        map.put("houston,tx", "IAH");
        map.put("phoenix,az", "PHX");
        map.put("philadelphia,pa", "PHL");
        map.put("san antonio,tx", "SAT");
        map.put("san diego,ca", "SAN");
        map.put("dallas,tx", "DFW");
        map.put("san jose,ca", "SJC");
        map.put("austin,tx", "AUS");
        map.put("jacksonville,fl", "JAX");
        map.put("san francisco,ca", "SFO");
        map.put("columbus,oh", "CMH");
        map.put("charlotte,nc", "CLT");
        map.put("indianapolis,in", "IND");
        map.put("seattle,wa", "SEA");
        map.put("denver,co", "DEN");
        map.put("nashville,tn", "BNA");
        map.put("oklahoma city,ok", "OKC");
        map.put("el paso,tx", "ELP");
        map.put("washington,dc", "DCA");
        map.put("las vegas,nv", "LAS");
        map.put("louisville,ky", "SDF");
        map.put("baltimore,md", "BWI");
        map.put("milwaukee,wi", "MKE");
        map.put("albuquerque,nm", "ABQ");
        map.put("tucson,az", "TUS");
        map.put("fresno,ca", "FAT");
        map.put("sacramento,ca", "SMF");
        map.put("kansas city,mo", "MCI");
        map.put("atlanta,ga", "ATL");
        map.put("miami,fl", "MIA");
        map.put("minneapolis,mn", "MSP");
        map.put("portland,or", "PDX");
        map.put("detroit,mi", "DTW");
        map.put("boston,ma", "BOS");
        map.put("memphis,tn", "MEM");
        map.put("new orleans,la", "MSY");
        map.put("cleveland,oh", "CLE");
        map.put("tampa,fl", "TPA");
        map.put("orlando,fl", "MCO");
        AIRPORTS = Collections.unmodifiableMap(map);
    }

    public static String lookup(String city, String state) throws IllegalArgumentException {
        String key = city.toLowerCase().trim() + "," + state.toLowerCase().trim();
        String code = AIRPORTS.get(key);
        if (code == null) {
            throw new IllegalArgumentException("no airport found for " + city + ", " + state);
        }
        return code;
    }
}
