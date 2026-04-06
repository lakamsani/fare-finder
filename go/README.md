# fare-finder — Go port

CLI tool to find the cheapest US domestic flight between two cities.

> Implements [issue #14](https://github.com/lakamsani/fare-finder/issues/14) — Go port.

## Prerequisites

- Go 1.22+
- A [SerpAPI](https://serpapi.com) key (free tier available)

## Build

```bash
cd go
go build -o fare-finder .
```

## Usage

```
./fare-finder <city1> <state1> <city2> <state2>
```

**Example:**

```bash
export SERPAPI_KEY=your_key_here
./fare-finder "San Francisco" CA "New York" NY
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
cd go
go test ./...
```

## Project Structure

```
go/
├── main.go                    # CLI entry point
├── go.mod
├── fare_finder/
│   ├── flight.go              # Flight type
│   ├── airport.go             # IATA airport lookup
│   └── searcher.go            # SerpAPI client + JSON parsing
└── tests/
    ├── airport_test.go
    └── searcher_test.go
```
