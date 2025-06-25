package main

func maxProduct(nums []int) int {
	ans, maxi, mini := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxi, mini = mini, maxi
		}
		mini = min(nums[i], mini*nums[i])
		maxi = max(nums[i], maxi*nums[i])
		ans = max(ans, maxi)
	}
	return ans
}
