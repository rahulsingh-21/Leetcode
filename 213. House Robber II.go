package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(houseRob(nums[:len(nums)-1]), houseRob(nums[1:]))
}

func houseRob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	fmt.Println(nums)
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
