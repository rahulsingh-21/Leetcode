package main

import "math"

func maxProfitIII(prices []int) int {

	stock1 := math.MaxInt32
	profit1 := 0
	stock2 := math.MaxInt32
	profit2 := 0

	for _, price := range prices {
		stock1 = min(stock1, price)
		profit1 = max(profit1, price-stock1)
		stock2 = min(stock2, price-profit1)
		profit2 = max(profit2, price-stock2)
	}

	return profit2
}
