package main

func deleteAndEarn(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum = max(sum, n)
	}

	freq := make([]int, sum+1)
	for _, n := range nums {
		freq[n] += n
	}

	dp := make([]int, sum+1)
	dp[0] = 0
	dp[1] = freq[1]
	for i := 2; i <= sum; i++ {
		dp[i] = max(dp[i-1], freq[i]+dp[i-2])
	}
	return dp[sum]
}

// 0 1 2 3 4
// 0 0 2 3 4

// 0 1 2 3 4
// 0 0 3 9 4
