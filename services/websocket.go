package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // Connected WebSocket clients
	broadcast = make(chan ProcessedOrderBook)  // Broadcast channel
	wsMutex   sync.Mutex                        // Mutex for thread safety
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WebSocketHandler upgrades HTTP connection to WebSocket
func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// Register client
	wsMutex.Lock()
	clients[conn] = true
	wsMutex.Unlock()

	log.Println("🔗 New WebSocket client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			wsMutex.Lock()
			delete(clients, conn)
			wsMutex.Unlock()
			break
		}
	}
}

// StartWebSocketServer starts the WebSocket server
func StartWebSocketServer() {
	go func() {
		for {
			data := <-OrderBookChannel // Get processed data
			wsMutex.Lock()
			for client := range clients {
				err := client.WriteJSON(data)
				if err != nil {
					log.Println("Error sending data:", err)
					client.Close()
					delete(clients, client)
				}
			}
			wsMutex.Unlock()
		}
	}()
}
