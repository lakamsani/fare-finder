# fare-finder

A CLI tool that finds the cheapest one-way flight between two US cities using the SerpAPI Google Flights engine.

## Build

```
go build -o fare-finder .
```

## Run

```
SERPAPI_KEY=<key> ./fare-finder "San Francisco" CA "New York" NY
```

## Test

```
go test ./...
```
