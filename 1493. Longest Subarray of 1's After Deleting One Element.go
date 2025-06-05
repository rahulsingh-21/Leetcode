package main

func longestSubarray(nums []int) int {
	i, j := 0, 0
	k := 1
	for j < len(nums) {
		if nums[j] == 0 {
			k--
		}
		j++
		if k < 0 {
			if nums[i] == 0 {
				k++
			}
			i++
		}
	}
	return j - i - 1
}
