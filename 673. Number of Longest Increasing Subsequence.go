package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	count := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}

	m, c := 0, 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = 0
				}

				if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}

			}

		}
	}

	fmt.Println(dp, count)
	for i := 0; i < n; i++ {
		m = max(m, dp[i])
	}

	for i := 0; i < n; i++ {
		if m == dp[i] {
			c += count[i]
		}
	}

	return c
}
