package services

import (
	"log"
	"sync"

	"github.com/shinwong97/models"

	"github.com/gorilla/websocket"
)

// OrderBookChannel to send processed data
var OrderBookChannel = make(chan models.ProcessedOrderBook, 10)

// WebSocketClient manages multiple WebSocket connections
type WebSocketClient struct {
	Connections map[string]*websocket.Conn
	mu          sync.Mutex
}

// NewWebSocketClient initializes the WebSocketClient
func NewWebSocketClient() *WebSocketClient {
	return &WebSocketClient{
		Connections: make(map[string]*websocket.Conn),
	}
}

// ConnectWebSocket connects to a given WebSocket URL dynamically
func ConnectWebSocket(exchange, wsURL string) {

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to %s WebSocket: %v", exchange, err)
	}
	defer conn.Close()

	log.Printf("Connected to %s WebSocket\n", exchange)

	for {
		// Read messages from WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from %s WebSocket: %v", exchange, err)
			break
		}

		// Log raw WebSocket data
		log.Printf("[%s] Received: %s\n", exchange, string(message))
		
	}
}


// CloseAll closes all WebSocket connections
func (w *WebSocketClient) CloseAll() {
	w.mu.Lock()
	defer w.mu.Unlock()
	for _, conn := range w.Connections {
		conn.Close()
	}
}
