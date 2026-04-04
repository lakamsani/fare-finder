# fare-finder

A CLI tool to find the cheapest flights between US cities using the [SerpAPI Google Flights engine](https://serpapi.com/google-flights-api).

## Requirements

- Java 17+
- Gradle 8+
- A [SerpAPI](https://serpapi.com) API key

## Usage

```bash
export SERPAPI_KEY=your_key_here
gradle run --args='"San Francisco" CA "New York" NY'
```

Or build a fat jar first:

```bash
gradle build
java -jar build/libs/fare-finder.jar "San Francisco" CA "New York" NY
```

Example output:

```
Searching for flights SFO → JFK on 2024-01-16...

🛫 Cheapest flight: SFO → JFK
✈️  JetBlue B6 123
💰 $189
🕗 Departs 2024-01-16 07:00 | Arrives 2024-01-16 15:30
⏱️  Duration: 5h 30m
```

## Running Tests

```bash
gradle test
```

## Supported Cities

Covers 42 major US cities including New York (JFK), Los Angeles (LAX), Chicago (ORD), San Francisco (SFO), and more.
