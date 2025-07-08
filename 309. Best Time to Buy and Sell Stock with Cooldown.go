package main

import "math"

func maxProfitV(prices []int) int {
	cool, sell, hold := 0, 0, math.MinInt32

	for i := 0; i < len(prices); i++ {
		prev_cool, prev_sell, prev_hold := cool, sell, hold

		cool = max(prev_cool, prev_sell)
		sell = prices[i] + prev_hold
		hold = max(prev_hold, prev_cool-prices[i])
	}

	return max(sell, cool)
}

//                  -> cooldown
// sell - > cooldown
//                  -> hold and buy - > hold -> sell
