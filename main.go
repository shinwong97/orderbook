package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shinwong97/services"
)

// func main() {
// 	services.OrderBook()
// }

// func main() {
// 	webSocketClient := services.NewWebSocketClient()

// 	// Connect to multiple WebSockets dynamically
// 	go webSocketClient.ConnectWebSocket("binance", "stream.binance.com:9443/ws/btcusdt@depth20")
// 	// go webSocketClient.ConnectWebSocket("example_exchange", "example.com/ws/orderbook")

// 	// Start REST API
// 	router := api.SetupRouter()
// 	log.Println("Starting API server on :8080")
// 	log.Fatal(router.Run(":8080"))
// }

// func main() {
// 	go services.ConnectWebSocket("binance", "wss://stream.binance.com:9443/ws/btcusdt@depth20")

// 	// Keep the main goroutine alive
// 	select {} // Prevents the main function from exiting
// }

func main() {
	r := gin.Default()

	// Define WebSocket route
	r.GET("/ws", services.WebSocketHandler)

	// Start WebSocket broadcasting
	go services.StartWebSocketServer()

	// Start multiple WebSocket connections
	exchanges := map[string]string{
		"binance": "wss://stream.binance.com:9443/ws/btcusdt@depth20",
		// Add more WebSocket connections here
	}
	go services.StartOrderBook(exchanges)

	log.Println("🚀 Server started on port 8080")
	r.Run(":8080")
}