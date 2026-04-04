# fare-finder

A CLI tool that searches for the cheapest US domestic flight between two cities using Google Flights data via SerpAPI.

> **Language:** Go 1.22+ (migrated from Python)

## Prerequisites

- Go 1.22+
- [SerpAPI](https://serpapi.com) account and API key

## Installation

```bash
git clone https://github.com/lakamsani/fare-finder.git
cd fare-finder
go build -o fare-finder .
```

## Usage

Set your SerpAPI key:

```bash
export SERPAPI_KEY="your_api_key_here"
```

Run the tool with origin and destination city/state:

```bash
./fare-finder "San Francisco" CA "New York" NY
./fare-finder "Chicago" IL "Miami" FL
./fare-finder "Seattle" WA "Austin" TX
./fare-finder "Boston" MA "Los Angeles" CA
```

### Sample Output

```
Searching for flights SFO → JFK on 2024-01-15...

🛫 Cheapest flight: SFO → JFK
✈️  United Airlines UA 101
💰 $189
🕗 Departs 2024-01-15 08:05 | Arrives 2024-01-15 16:23
⏱️  Duration: 5h 18m
```

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `SERPAPI_KEY` | Yes | Your SerpAPI API key. Get one at https://serpapi.com |

## Supported Cities

40+ major US cities including New York, Los Angeles, Chicago, Houston, Phoenix, Philadelphia, San Francisco, Seattle, Denver, Atlanta, Miami, Boston, Dallas, and more.

## Running Tests

```bash
go test -v ./...
```

## Project Structure

```
fare-finder/
├── main.go               # CLI entrypoint
├── airports/
│   ├── airports.go       # City → IATA code lookup
│   └── airports_test.go
├── flights/
│   ├── flights.go        # SerpAPI search + parsing
│   └── flights_test.go
└── go.mod
```

## SerpAPI Free Tier

SerpAPI offers a free tier with 100 searches/month. Each fare-finder call uses one API call. See [SerpAPI pricing](https://serpapi.com/pricing) for higher usage.

## License

MIT
