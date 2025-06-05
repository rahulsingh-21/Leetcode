package main

func longestOnes(nums []int, k int) int {
	i := 0
	j := 0
	for i = 0; i < len(nums); i++ {

		if nums[i] == 0 {
			k--
		}

		if k < 0 {
			if nums[j] == 0 {
				k++
			}
			j++
		}
	}
	return i - j
}
