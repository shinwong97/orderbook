# Order Book WebSocket Client

This project connects to Binance's WebSocket API to fetch the BTC/USDT order book and calculates the weighted average price based on the bids and asks.

## Prerequisites

Ensure you have the following installed:
- [Go](https://go.dev/doc/install) (1.18 or later)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/shinwong97/orderbook.git
   cd your-repo
   ```

2. Initialize the Go module:
   ```sh
   go mod init your-repo
   go mod tidy
   ```

## Running the Program

1. Build and run the program:
   ```sh
   go run main.go
   ```

   If the project has multiple files:
   ```sh
   go run .
   ```

2. The output will display the average price and order book details in the console.

## Project Structure

```
.
├── main.go          # Entry point of the program
├── orderbook.go     # Handles WebSocket connection and order book processing
├── utils            # Utility functions (e.g., ParseFloat)
│   ├── utils.go
├── go.mod          # Go module file
├── go.sum          # Dependencies checksum
└── README.md       # This file
```

## Configuration

Modify `OrderBook()` in `orderbook.go` to change the trading pair:
```go
Path: "/ws/ethusdt@depth20",
```
For Ethereum instead of Bitcoin.

## Dependencies

This project uses the following dependencies:
```sh
go get -u github.com/gorilla/websocket
```

## License

This project is licensed under the MIT License.

