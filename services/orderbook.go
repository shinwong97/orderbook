package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/shinwong97/utils"
)

type OrderBookUpdate struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

func OrderBook() {
	// Binance WebSocket URL for BTCUSDT order book
	webSocketURL := url.URL{
		Scheme: "wss",
		Host:   "stream.binance.com:9443",
		Path:   "/ws/btcusdt@depth20", // depth20 gives us 20 levels
	}

	// Connect to WebSocket
	connectWebSocket, _, err := websocket.DefaultDialer.Dial(webSocketURL.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connectWebSocket.Close()

	fmt.Println("Connected to Binance WebSocket")

	for {
		// Read message
		_, message, err := connectWebSocket.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		// Parse order book update
		var orderBook OrderBookUpdate
		if err := json.Unmarshal(message, &orderBook); err != nil {
			log.Println("Error unmarshalling JSON::", err)
			continue
		}

		// Compute the weighted average price
		averagePrice, totalOrders, totalSize := calculateWeightedAverage(orderBook)

		fmt.Printf("📈 Average Price: %.2f USDT\n", averagePrice)
		fmt.Printf("📊 Total Orders: %d\n", totalOrders)
		fmt.Printf("📦 Total Size: %.8f BTC\n\n", totalSize)
	}
}


func calculateWeightedAverage(orderBook OrderBookUpdate) (float64, int, float64) {
	var totalWeightedPrice float64
	var totalSize float64
	var totalOrders int

	// Process bids
	for _, bid := range orderBook.Bids {
		price := utils.ParseFloat(bid[0])
		size := utils.ParseFloat(bid[1])
		totalWeightedPrice += price * size
		totalSize += size
		totalOrders++
	}

	// Process asks
	for _, ask := range orderBook.Asks {
		price := utils.ParseFloat(ask[0])
		size := utils.ParseFloat(ask[1])
		totalWeightedPrice += price * size
		totalSize += size
		totalOrders++
	}

	// Prevent division by zero
	if totalSize == 0 {
		return 0, totalOrders, totalSize
	}

	// Calculate weighted average price
	averagePrice := totalWeightedPrice / totalSize
	return averagePrice, totalOrders, totalSize
}