# fare-finder

A command-line tool for finding the cheapest one-way flights between US cities using the SerpAPI Google Flights engine.

## Requirements

- Go 1.21+
- A [SerpAPI](https://serpapi.com) API key

## Run

```bash
export SERPAPI_KEY=your_key_here
cd go
go run . "San Francisco" CA "New York" NY
```

Or build first:

```bash
cd go
go build -o fare-finder .
./fare-finder "San Francisco" CA "New York" NY
```

## Test

```bash
cd go
go test -v ./...
```

## Supported Cities

Covers 42 major US cities including New York (JFK), Los Angeles (LAX), Chicago (ORD), San Francisco (SFO), and more.
