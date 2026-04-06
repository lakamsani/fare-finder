# fare-finder (Kotlin)

Kotlin port of fare-finder using Jackson for JSON parsing and the Java HTTP client.

## Prerequisites

- JDK 11+
- Gradle 8+ (or use the wrapper if added)
- A [SerpAPI](https://serpapi.com) key

## Build

```bash
cd kotlin
gradle build
```

## Run

```bash
export SERPAPI_KEY=your_key_here
gradle run --args='"San Francisco" CA "New York" NY'
```

Or build a fat JAR and run directly:

```bash
gradle jar
java -jar build/libs/fare-finder-kotlin-1.0.0.jar "San Francisco" CA "New York" NY
```

## Test

```bash
gradle test
```
