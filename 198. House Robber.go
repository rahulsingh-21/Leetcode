package main

func robI(nums []int) int {
	dp := make([]int, len(nums))
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1]
}
