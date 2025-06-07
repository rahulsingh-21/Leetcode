package main

func maximumProfit(prices []int, k int) int64 {

	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	prev := make([][3]int, k+1)
	curr := make([][3]int, k+1)

	for i := 0; i <= k; i++ {
		prev[i][1] = -prices[0]
		prev[i][2] = prices[0]
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			curr[j][0] = prev[j][0]
			if j > 0 {
				curr[j][0] = max(curr[j][0], prev[j-1][1]+prices[i])
				curr[j][0] = max(curr[j][0], prev[j-1][2]-prices[i])
			}

			curr[j][1] = prev[j][1]
			if prev[j][0]-prices[i] > curr[j][1] {
				curr[j][1] = prev[j][0] - prices[i]
			}

			curr[j][2] = prev[j][2]
			if prev[j][0]+prices[i] > curr[j][2] {
				curr[j][2] = prev[j][0] + prices[i]
			}

		}
		prev, curr = curr, prev
	}
	var ans int64
	for i := 0; i <= k; i++ {
		//ans = max(ans, int64(prev[i][0]))  check why its giving error on vscode
	}
	return ans
}
