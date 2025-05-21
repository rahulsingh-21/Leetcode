package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {

	sum := 0
	for _, n := range nums {
		sum += n
	}

	t := (sum + target) / 2

	if (sum+target)%2 != 0 || sum < abs(target) {
		return 0
	}

	dp := make([]int, t+1)
	dp[0] = 1
	fmt.Println(t)
	for _, num := range nums {
		for j := t; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[t]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
