# fare-finder — Rust port

CLI tool to find the cheapest US domestic flight between two cities.

> Implements [issue #17](https://github.com/lakamsani/fare-finder/issues/17) — Rust port.

## Prerequisites

- Rust 1.75+ (via [rustup](https://rustup.rs))
- A [SerpAPI](https://serpapi.com) key (free tier available)

## Build

```bash
cd rust
cargo build --release
```

## Usage

```
./target/release/fare-finder <city1> <state1> <city2> <state2>
```

**Example:**

```bash
export SERPAPI_KEY=your_key_here
./target/release/fare-finder "San Francisco" CA "New York" NY
```

**Output:**
```
Searching for flights SFO -> JFK on 2024-01-16...

Cheapest flight: SFO -> JFK
United Airlines UA 101
$189
Departs 2024-01-16 08:05 | Arrives 2024-01-16 16:23
Duration: 5h 18m
```

## Run Tests

```bash
cd rust
cargo test
```

## Project Structure

```
rust/
├── Cargo.toml
├── src/
│   ├── lib.rs          # Public library surface
│   ├── main.rs         # CLI entry point
│   ├── flight.rs       # Flight struct
│   ├── airport.rs      # IATA airport lookup
│   └── searcher.rs     # SerpAPI client + JSON parsing
└── tests/
    ├── airport_test.rs
    └── searcher_test.rs
```
