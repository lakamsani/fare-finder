# fare-finder

A command-line tool for finding the cheapest one-way flights between US cities using the SerpAPI Google Flights engine.

## Go

### Build

```
go build -o fare-finder .
```

### Run

```
SERPAPI_KEY=<key> ./fare-finder "San Francisco" CA "New York" NY
```

### Test

```
go test ./...
```

## Python

### Install

```
pip install .
```

### Run

```
SERPAPI_KEY=<key> python -m fare_finder "San Francisco" CA "New York" NY
```

### Test

```
pip install pytest
pytest
```
