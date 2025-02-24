package services

import (
	"encoding/json"
	"log"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/shinwong97/utils"
)

type OrderBookUpdate struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type ProcessedOrderBook struct {
	Exchange     string  `json:"exchange"`
	AveragePrice float64 `json:"average_price"`
	TotalOrders  int     `json:"total_orders"`
	TotalSize    float64 `json:"total_size"`
}

var (
	OrderBookChannel = make(chan ProcessedOrderBook, 10) // Channel to send processed data
	wg               sync.WaitGroup                      // WaitGroup for concurrency
)

// StartOrderBook initializes multiple WebSocket connections dynamically
func StartOrderBook(exchanges map[string]string) {
	for exchange, wsURL := range exchanges {
		wg.Add(1)
		go connectToExchange(exchange, wsURL)
	}
	wg.Wait() // Wait for all goroutines to finish
}

func connectToExchange(exchange, wsURL string) {
	defer wg.Done()

	parsedURL, err := url.Parse(wsURL)
	if err != nil {
		log.Fatalf("Invalid WebSocket URL for %s: %v", exchange, err)
	}

	// Connect to WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(parsedURL.String(), nil)
	if err != nil {
		log.Printf("Failed to connect to %s WebSocket: %v", exchange, err)
		return
	}
	defer conn.Close()

	log.Printf("✅ Connected to %s WebSocket", exchange)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from %s: %v", exchange, err)
			return
		}

		// Parse order book data
		var orderBook OrderBookUpdate
		if err := json.Unmarshal(message, &orderBook); err != nil {
			log.Println("Error unmarshalling JSON:", err)
			continue
		}

		// Compute weighted average price
		averagePrice, totalOrders, totalSize := calculateWeightedAverage(orderBook)

		// Send processed data to channel
		OrderBookChannel <- ProcessedOrderBook{
			Exchange:     exchange,
			AveragePrice: averagePrice,
			TotalOrders:  totalOrders,
			TotalSize:    totalSize,
		}
	}
}

// Calculate weighted average price
func calculateWeightedAverage(orderBook OrderBookUpdate) (float64, int, float64) {
	var totalWeightedPrice, totalSize float64
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
