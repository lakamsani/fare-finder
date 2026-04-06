use fare_finder::searcher::parse_flights;

const MOCK_JSON: &str = r#"{
    "best_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 08:05"},
                    "arrival_airport":   {"time": "2024-01-15 16:23"},
                    "duration": 318,
                    "airline": "United Airlines",
                    "flight_number": "UA 101"
                }
            ],
            "price": 189
        }
    ],
    "other_flights": [
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 10:30"},
                    "arrival_airport":   {"time": "2024-01-15 18:45"},
                    "duration": 315,
                    "airline": "Delta Air Lines",
                    "flight_number": "DL 405"
                }
            ],
            "price": 245
        },
        {
            "flights": [
                {
                    "departure_airport": {"time": "2024-01-15 06:00"},
                    "arrival_airport":   {"time": "2024-01-15 14:10"},
                    "duration": 310,
                    "airline": "JetBlue",
                    "flight_number": "B6 816"
                }
            ],
            "price": 159
        }
    ]
}"#;

#[test]
fn test_parse_flights_sorted_by_price() {
    let flights = parse_flights(MOCK_JSON).expect("parse_flights failed");

    assert_eq!(flights.len(), 3, "expected 3 flights, got {}", flights.len());

    // Sorted by price ascending: 159, 189, 245
    assert_eq!(flights[0].price, 159);
    assert_eq!(flights[1].price, 189);
    assert_eq!(flights[2].price, 245);

    // Cheapest is JetBlue B6 816
    assert_eq!(flights[0].airline, "JetBlue");
    assert_eq!(flights[0].flight_number, "B6 816");
    assert_eq!(flights[0].duration, "5h 10m");
}

#[test]
fn test_parse_flights_empty_json() {
    let flights = parse_flights("{}").expect("parse_flights failed on empty object");
    assert!(flights.is_empty(), "expected empty list for empty JSON");
}

#[test]
fn test_parse_flights_invalid_json() {
    let result = parse_flights("not json at all");
    assert!(result.is_err(), "expected error for invalid JSON");
    let msg = result.unwrap_err();
    assert!(msg.contains("failed to parse JSON"), "unexpected error: {}", msg);
}

#[test]
fn test_missing_serpapi_key_returns_error() {
    // Ensure the env var is unset for this test
    std::env::remove_var("SERPAPI_KEY");
    // Also clear SERPAPI_BASE_URL so we don't hit a real server
    std::env::remove_var("SERPAPI_BASE_URL");

    let result = fare_finder::searcher::search_flights("SFO", "JFK", "2024-01-15");
    assert!(result.is_err(), "expected error when SERPAPI_KEY is missing");
    let msg = result.unwrap_err();
    assert!(
        msg.contains("SERPAPI_KEY"),
        "expected 'SERPAPI_KEY' in error message, got: {:?}",
        msg
    );
}
