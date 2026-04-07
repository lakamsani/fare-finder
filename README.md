# fare-finder

A command-line tool for finding the cheapest one-way flights between US cities using the SerpAPI Google Flights engine.

## Requirements

- Node.js 18+
- A [SerpAPI](https://serpapi.com) API key

## Install

```bash
npm install
```

## Run

```bash
export SERPAPI_KEY=your_key_here
npm run build
npm start -- "San Francisco" CA "New York" NY
```

## Test

```bash
npm test
```

## Supported Cities

Covers 42 major US cities including New York (JFK), Los Angeles (LAX), Chicago (ORD), San Francisco (SFO), and more.
