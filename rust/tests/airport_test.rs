use fare_finder::airport::lookup_airport;

#[test]
fn test_known_airports() {
    let cases = vec![
        ("San Francisco", "CA", "SFO"),
        ("New York",      "NY", "JFK"),
        ("Chicago",       "IL", "ORD"),
        ("Atlanta",       "GA", "ATL"),
        ("Denver",        "CO", "DEN"),
        ("Seattle",       "WA", "SEA"),
        ("Miami",         "FL", "MIA"),
        ("Las Vegas",     "NV", "LAS"),
        ("Boston",        "MA", "BOS"),
        ("Dallas",        "TX", "DFW"),
    ];

    for (city, state, want) in cases {
        let got = lookup_airport(city, state)
            .unwrap_or_else(|e| panic!("lookup_airport({:?}, {:?}) failed: {}", city, state, e));
        assert_eq!(got, want, "lookup_airport({:?}, {:?})", city, state);
    }
}

#[test]
fn test_unknown_city_returns_error() {
    let result = lookup_airport("Nonexistent City", "XX");
    assert!(result.is_err(), "expected Err for unknown city, got Ok");
    let msg = result.unwrap_err();
    assert!(!msg.is_empty(), "error message should not be empty");
}
