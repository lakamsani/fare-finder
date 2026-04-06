use std::collections::HashMap;

/// Returns a map from "city,state" (lowercase) to IATA airport code.
fn airports() -> HashMap<&'static str, &'static str> {
    let mut m = HashMap::new();
    m.insert("new york,ny",      "JFK");
    m.insert("los angeles,ca",   "LAX");
    m.insert("chicago,il",       "ORD");
    m.insert("houston,tx",       "IAH");
    m.insert("phoenix,az",       "PHX");
    m.insert("philadelphia,pa",  "PHL");
    m.insert("san antonio,tx",   "SAT");
    m.insert("san diego,ca",     "SAN");
    m.insert("dallas,tx",        "DFW");
    m.insert("san jose,ca",      "SJC");
    m.insert("austin,tx",        "AUS");
    m.insert("jacksonville,fl",  "JAX");
    m.insert("san francisco,ca", "SFO");
    m.insert("columbus,oh",      "CMH");
    m.insert("charlotte,nc",     "CLT");
    m.insert("indianapolis,in",  "IND");
    m.insert("seattle,wa",       "SEA");
    m.insert("denver,co",        "DEN");
    m.insert("nashville,tn",     "BNA");
    m.insert("oklahoma city,ok", "OKC");
    m.insert("el paso,tx",       "ELP");
    m.insert("washington,dc",    "DCA");
    m.insert("las vegas,nv",     "LAS");
    m.insert("louisville,ky",    "SDF");
    m.insert("baltimore,md",     "BWI");
    m.insert("milwaukee,wi",     "MKE");
    m.insert("albuquerque,nm",   "ABQ");
    m.insert("tucson,az",        "TUS");
    m.insert("fresno,ca",        "FAT");
    m.insert("sacramento,ca",    "SMF");
    m.insert("kansas city,mo",   "MCI");
    m.insert("atlanta,ga",       "ATL");
    m.insert("miami,fl",         "MIA");
    m.insert("minneapolis,mn",   "MSP");
    m.insert("portland,or",      "PDX");
    m.insert("detroit,mi",       "DTW");
    m.insert("boston,ma",        "BOS");
    m.insert("memphis,tn",       "MEM");
    m.insert("new orleans,la",   "MSY");
    m.insert("cleveland,oh",     "CLE");
    m.insert("tampa,fl",         "TPA");
    m.insert("orlando,fl",       "MCO");
    m
}

/// Returns the IATA airport code for a given city and state.
///
/// # Errors
/// Returns an error string if no airport is found for the city/state pair.
pub fn lookup_airport(city: &str, state: &str) -> Result<&'static str, String> {
    let key = format!("{},{}", city.trim().to_lowercase(), state.trim().to_lowercase());
    airports()
        .get(key.as_str())
        .copied()
        .ok_or_else(|| format!("no airport found for {}, {}", city, state))
}
