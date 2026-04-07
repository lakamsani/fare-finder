# fare-finder

A command-line tool for finding the cheapest one-way flights between US cities using the SerpAPI Google Flights engine.

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
