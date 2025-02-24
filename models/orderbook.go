package models

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
