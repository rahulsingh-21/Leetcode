package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	t := sum / 2

	dp := make([]bool, t+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for j := t; j >= n; j-- {
			if dp[j-n] {
				dp[j] = true
			}
		}
	}
	return dp[t]
}
