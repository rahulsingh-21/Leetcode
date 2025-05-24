package main

func maxProfitIV(k int, prices []int) int {

	n := len(prices)

	if n == 0 || k == 0 {
		return 0
	}

	prev := make([]int, n)
	curr := make([]int, n)

	for i := 1; i <= k; i++ {
		maxDiff := -prices[0]
		for j := 1; j < n; j++ {
			curr[j] = max(curr[j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, prev[j]-prices[j])
		}
		prev, curr = curr, prev
	}

	return prev[n-1]
}
