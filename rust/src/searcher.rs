use serde::Deserialize;
use std::env;

use crate::flight::Flight;

/// SerpAPI endpoint — overridable via the `SERPAPI_BASE_URL` env var (for tests).
fn base_url() -> String {
    env::var("SERPAPI_BASE_URL").unwrap_or_else(|_| "https://serpapi.com/search".to_string())
}

// ---------- JSON shapes ----------

#[derive(Deserialize)]
struct SerpResponse {
    #[serde(default)]
    best_flights: Vec<FlightGroup>,
    #[serde(default)]
    other_flights: Vec<FlightGroup>,
}

#[derive(Deserialize)]
struct FlightGroup {
    #[serde(default)]
    flights: Vec<FlightLeg>,
    #[serde(default)]
    price: u32,
}

#[derive(Deserialize)]
struct FlightLeg {
    #[serde(default)]
    airline: String,
    #[serde(default)]
    flight_number: String,
    #[serde(default)]
    duration: u32,
    #[serde(default)]
    departure_airport: TimeNode,
    #[serde(default)]
    arrival_airport: TimeNode,
}

#[derive(Deserialize, Default)]
struct TimeNode {
    #[serde(default)]
    time: String,
}

// ---------- Public API ----------

/// Fetches flights from SerpAPI for the given origin, destination, and date.
///
/// * `origin` / `dest` — 3-letter IATA codes  
/// * `date` — `YYYY-MM-DD`
///
/// Returns flights sorted by price ascending.
///
/// # Errors
/// Returns an error string if the API key is missing, the HTTP request fails,
/// or the response is not valid JSON.
pub fn search_flights(origin: &str, dest: &str, date: &str) -> Result<Vec<Flight>, String> {
    let api_key = env::var("SERPAPI_KEY")
        .unwrap_or_default();

    if api_key.is_empty() {
        return Err(
            "SERPAPI_KEY environment variable is not set. Get a key at https://serpapi.com"
                .to_string(),
        );
    }

    let url = format!(
        "{}?engine=google_flights&departure_id={}&arrival_id={}&outbound_date={}\
         &currency=USD&hl=en&api_key={}&type=2",
        base_url(),
        origin,
        dest,
        date,
        api_key,
    );

    let response = ureq::get(&url)
        .call()
        .map_err(|e| format!("HTTP request failed: {}", e))?;

    if response.status() != 200 {
        return Err(format!("API returned status {}", response.status()));
    }

    let body = response
        .into_string()
        .map_err(|e| format!("failed to read response body: {}", e))?;

    parse_flights(&body)
}

/// Parses a SerpAPI JSON response and returns flights sorted by price ascending.
///
/// # Errors
/// Returns an error string if the JSON is malformed.
pub fn parse_flights(json: &str) -> Result<Vec<Flight>, String> {
    let data: SerpResponse =
        serde_json::from_str(json).map_err(|e| format!("failed to parse JSON: {}", e))?;

    let mut flights = Vec::new();

    for group in data.best_flights.iter().chain(data.other_flights.iter()) {
        if group.flights.is_empty() {
            continue;
        }
        let first = &group.flights[0];
        let last = &group.flights[group.flights.len() - 1];
        let total_duration: u32 = group.flights.iter().map(|l| l.duration).sum();

        flights.push(Flight {
            airline: first.airline.clone(),
            flight_number: first.flight_number.clone(),
            price: group.price,
            departure_time: first.departure_airport.time.clone(),
            arrival_time: last.arrival_airport.time.clone(),
            duration: format_duration(total_duration),
        });
    }

    flights.sort_by_key(|f| f.price);
    Ok(flights)
}

fn format_duration(minutes: u32) -> String {
    format!("{}h {}m", minutes / 60, minutes % 60)
}
