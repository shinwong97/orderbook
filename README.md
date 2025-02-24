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

## Running the WebSocket Server

### **1. Start the Server**

Run the WebSocket server locally:

```sh
go run main.go
```

You should see:

```
🚀 Server started on port 8080
```

---

## Testing the WebSocket API

### **Using `wscat` (WebSocket Client)**

#### **1. Install `wscat` (Optional)**

If you don’t have `wscat` installed, install it using npm:

```sh
npm install -g wscat
```

#### **2. Connect to the WebSocket**

If your server is running **locally**, use:

```sh
wscat -c ws://localhost:8080/ws
```

If your server is **running on a remote server**, use:

```sh
wscat -c ws://your-server-ip:8080/ws
```

Replace `your-server-ip` with your actual server IP or domain.

#### **3. Expected Output**

Once connected, you should receive real-time order book updates like:

```json
{
  "exchange": "binance",
  "average_price": 43250.12,
  "total_orders": 40,
  "total_size": 105.4
}
```



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

