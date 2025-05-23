package main

func maxProfitII(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			ans += prices[i+1] - prices[i]
		}
	}
	return ans
}
