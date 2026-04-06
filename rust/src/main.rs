// fare-finder: CLI tool to find the cheapest US domestic flight between two cities (Rust port).
//
// Fixes https://github.com/lakamsani/fare-finder/issues/17

use fare_finder::airport;
use fare_finder::searcher;

use std::process;

fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() != 5 {
        eprintln!("Usage: fare-finder <city1> <state1> <city2> <state2>");
        eprintln!(r#"Example: fare-finder "San Francisco" CA "New York" NY"#);
        process::exit(1);
    }

    let city1 = title_case(args[1].trim());
    let state1 = args[2].trim().to_uppercase();
    let city2 = title_case(args[3].trim());
    let state2 = args[4].trim().to_uppercase();

    let origin = match airport::lookup_airport(&city1, &state1) {
        Ok(code) => code,
        Err(e) => {
            eprintln!("Error: {}", e);
            process::exit(1);
        }
    };

    let dest = match airport::lookup_airport(&city2, &state2) {
        Ok(code) => code,
        Err(e) => {
            eprintln!("Error: {}", e);
            process::exit(1);
        }
    };

    let tomorrow = {
        // chrono not used to keep deps minimal — compute date manually via std
        let now = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .expect("time went backwards")
            .as_secs();
        let days = now / 86400 + 1; // days since epoch + 1 = tomorrow
        epoch_days_to_date(days)
    };

    println!("Searching for flights {} -> {} on {}...\n", origin, dest, tomorrow);

    let flights = match searcher::search_flights(origin, dest, &tomorrow) {
        Ok(f) => f,
        Err(e) => {
            eprintln!("Error: {}", e);
            process::exit(1);
        }
    };

    if flights.is_empty() {
        println!("No flights found for this route and date. Try a different date or city pair.");
        process::exit(0);
    }

    let cheapest = &flights[0];
    println!("Cheapest flight: {} -> {}", origin, dest);
    println!("{} {}", cheapest.airline, cheapest.flight_number);
    println!("${}", cheapest.price);
    println!("Departs {} | Arrives {}", cheapest.departure_time, cheapest.arrival_time);
    println!("Duration: {}", cheapest.duration);
}

/// Converts a string so that each word starts with an uppercase letter
/// and the remaining letters are lowercase.
fn title_case(s: &str) -> String {
    s.split_whitespace()
        .map(|word| {
            let mut chars = word.chars();
            match chars.next() {
                None => String::new(),
                Some(first) => {
                    first.to_uppercase().to_string()
                        + &chars.as_str().to_lowercase()
                }
            }
        })
        .collect::<Vec<_>>()
        .join(" ")
}

/// Converts days since Unix epoch to a "YYYY-MM-DD" string without external crates.
fn epoch_days_to_date(days: u64) -> String {
    // Algorithm: civil date from days since epoch (1970-01-01)
    // Reference: http://howardhinnant.github.io/date_algorithms.html
    let z = days as i64 + 719468;
    let era = if z >= 0 { z } else { z - 146096 } / 146097;
    let doe = (z - era * 146097) as u64;
    let yoe = (doe - doe / 1460 + doe / 36524 - doe / 146096) / 365;
    let y = yoe as i64 + era * 400;
    let doy = doe - (365 * yoe + yoe / 4 - yoe / 100);
    let mp = (5 * doy + 2) / 153;
    let d = doy - (153 * mp + 2) / 5 + 1;
    let m = if mp < 10 { mp + 3 } else { mp - 9 };
    let y = if m <= 2 { y + 1 } else { y };
    format!("{:04}-{:02}-{:02}", y, m, d)
}
