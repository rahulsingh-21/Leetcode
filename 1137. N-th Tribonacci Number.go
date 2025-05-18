package main

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n-1]
}
