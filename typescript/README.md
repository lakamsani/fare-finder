# fare-finder-ts

TypeScript/Node.js port of the fare-finder service.

## Usage

```bash
npm install
npm run build
node dist/src/main.js "San Francisco" CA "New York" NY
```

Or run directly with ts-node:

```bash
npm start -- "San Francisco" CA "New York" NY
```

Set the `SERPAPI_KEY` environment variable before running:

```bash
export SERPAPI_KEY=your_key_here
```

## Tests

```bash
npm test
```
