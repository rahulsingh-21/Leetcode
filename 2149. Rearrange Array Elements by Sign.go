package main

func rearrangeArray(nums []int) []int {
	ans := make([]int, len(nums))

	i := 0
	j := 1

	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			ans[i] = nums[k]
			i += 2
		} else {
			ans[j] = nums[k]
			j += 2
		}
	}
	return ans
}
