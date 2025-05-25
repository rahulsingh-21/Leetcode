package main

func lengthOfLIS(nums []int) int {
	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := 0
	for _, n := range dp {
		ans = max(ans, n)
	}

	return ans
}
