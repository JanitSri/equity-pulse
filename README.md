
# Equity Pulse 📈 💰 

Equity Pulse is a CLI application designed to provide quick and easy access to stock market information. With Equity Pulse, you can fetch details about specific stocks, see which stocks are trending, and view upcoming calendar events related to the stock market.

## TOC
- [Commands](#commands)
  - [Get Stock Information](#get-stock-information)
  - [Trending Stocks](#trending-stocks)
  - [Calendar Events](#calendar-events)
- [Examples](#examples)
- [Build \& Run](#build--run)
  - [Redis \& Redis Insights](#redis--redis-insights)
  - [Equity Pulse](#equity-pulse)

## Commands

### Get Stock Information

Retrieve information about a specific stock by its ticker symbol or company name.

- By Ticker Symbol:
  ```bash
  equity-pulse stock --ticker AAPL
  ```

  ```bash
  equity-pulse stock --name Apple
  ```

### Trending Stocks

View a list of currently trending stocks.

```bash
equity-pulse trending
```

### Calendar Events

View upcoming stock market-related events within a specified date range.

```bash
equity-pulse events --startdate YYYY-MM-DD --enddate YYYY-MM-DD
```

- **Flags:**
  - `startdate` - The start date for the events in YYYY-MM-DD format.
  - `enddate` - The end date for the events in YYYY-MM-DD format.

## Examples

- Fetch stock information for Apple Inc. by ticker symbol:
  ```bash
  equity-pulse stock --ticker AAPL
  ```

- Fetch stock information for Apple Inc. by company name:
  ```bash
  equity-pulse stock --name Apple
  ```

- Show trending stocks:
  ```bash
  equity-pulse trending
  ```

- Show calendar events from June 1, 2024, to June 30, 2024:
  ```bash
  equity-pulse events --startdate 2024-06-01 --enddate 2024-06-30

## Build & Run 

### Redis & Redis Insights 
```bash
  make compose-up
```

### Equity Pulse
```bash
  make run 
```