package main

func minSubArrayLen(target int, nums []int) int {
	l, r, sum, ans := 0, 0, 0, len(nums)+1
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
