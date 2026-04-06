# fare-finder (Go)

A CLI tool to find the cheapest flights between US cities using the
[SerpAPI Google Flights engine](https://serpapi.com/google-flights-api).

## Requirements

- Go 1.21+
- A [SerpAPI](https://serpapi.com) API key

## Setup

```bash
cd go
go build -o fare-finder .
```

## Usage

```bash
export SERPAPI_KEY=your_key_here
./fare-finder "San Francisco" CA "New York" NY
```

Or run directly:

```bash
go run . "San Francisco" CA "New York" NY
```

Example output:

```
Searching for flights SFO -> JFK on 2024-01-16...

Cheapest flight: SFO -> JFK
JetBlue B6 123
$189
Departs 2024-01-16 07:00 | Arrives 2024-01-16 15:30
Duration: 5h 30m
```

## Running Tests

```bash
cd go
go test ./...
```

## Supported Cities

Covers 42 major US cities including New York (JFK), Los Angeles (LAX),
Chicago (ORD), San Francisco (SFO), and more.
