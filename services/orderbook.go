package services

import (
	"github.com/shinwong97/models"

	"github.com/shinwong97/utils"

	"fmt"
)

//  computes the weighted average price from order book data
func CalculateWeightedAverage(orderBook models.OrderBookUpdate) (float64, int, float64) {
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

	if totalSize == 0 {
		return 0, totalOrders, totalSize
	}

	averagePrice := totalWeightedPrice / totalSize

		fmt.Printf("Average Price: %.2f USDT\n", averagePrice)
		fmt.Printf(" Total Orders: %d\n", totalOrders)
		fmt.Printf(" Total Size: %.8f BTC\n\n", totalSize)

	return averagePrice, totalOrders, totalSize
}
