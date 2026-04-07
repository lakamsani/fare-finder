# fare-finder

A command-line tool for finding the cheapest one-way flights between US cities using the SerpAPI Google Flights engine.

## Build

```
mvn package
```

## Run

```
SERPAPI_KEY=<key> java -jar target/fare-finder-1.0-SNAPSHOT.jar "San Francisco" CA "New York" NY
```

## Test

```
mvn test
```
