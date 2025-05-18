package main

func fib(n int) int {
	dp := make([]int, n)
	if n < 2 {
		return n
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n-1]
}
