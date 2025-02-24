package utils

import (
	"fmt"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ParseFloat(value string) float64 {
	floatedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return 0 
	}
	return floatedValue
}


// Upgrader handles WebSocket upgrade
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (CORS)
	},
}

// UpgradeConnection upgrades HTTP to WebSocket
func UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return nil, err
	}
	return conn, nil
}