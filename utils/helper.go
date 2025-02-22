package utils

import (
	"fmt"
	"strconv"
)

func ParseFloat(value string) float64 {
	floatedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return 0 
	}
	return floatedValue
}
