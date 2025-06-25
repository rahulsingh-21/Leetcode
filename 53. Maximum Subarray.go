package main

func maxSubArray(nums []int) int {
	sum, ans := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		ans = max(ans, sum)
	}
	return ans
}
