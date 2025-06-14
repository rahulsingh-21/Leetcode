package main

func sortArrayByParity(nums []int) []int {
	ans := make([]int, len(nums))

	l, r := 0, len(nums)-1
	i := 0
	for l <= r {
		if nums[i]%2 == 0 {
			ans[l] = nums[i]
			l++
		} else {
			ans[r] = nums[i]
			r--
		}
		i++
	}
	return ans
}
