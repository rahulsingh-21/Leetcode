package main

import "math"

func maxProfit(prices []int) int {
	m := 0
	curr := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		curr = min(curr, prices[i])
		m = max(m, prices[i]-curr)
	}
	return m
}
