package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinwong97/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/orderbook",GetOrderBook)

	return router
}

func GetOrderBook(c *gin.Context) {
	select {
	case orderBook := <-services.OrderBookChannel:
		c.JSON(http.StatusOK, orderBook)
	default:
		c.JSON(http.StatusNoContent, gin.H{"message": "No new order book data available"})
	}
}
