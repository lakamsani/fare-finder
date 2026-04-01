# fare-finder

A Python CLI tool that searches for the cheapest US domestic flight between two cities using Google Flights data via SerpAPI.

## Prerequisites

- Python 3.10+
- [SerpAPI](https://serpapi.com) account and API key

## Installation

```bash
git clone https://github.com/lakamsani/fare-finder.git
cd fare-finder
pip install -r requirements.txt
```

## Usage

Set your SerpAPI key:

```bash
export SERPAPI_KEY="your_api_key_here"
```

Run the tool with origin and destination city/state:

```bash
python fare_finder.py "San Francisco" CA "New York" NY
python fare_finder.py "Chicago" IL "Miami" FL
python fare_finder.py "Seattle" WA "Austin" TX
python fare_finder.py "Boston" MA "Los Angeles" CA
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

The tool supports 40+ major US cities including New York, Los Angeles, Chicago, Houston, Phoenix, Philadelphia, San Francisco, Seattle, Denver, Atlanta, Miami, Boston, Dallas, and many more.

## Running Tests

```bash
pip install pytest
pytest tests/
```

## SerpAPI Free Tier

SerpAPI offers a free tier with 100 searches per month. Each fare-finder search uses one API call. For higher usage, see [SerpAPI pricing](https://serpapi.com/pricing).

## License

MIT
