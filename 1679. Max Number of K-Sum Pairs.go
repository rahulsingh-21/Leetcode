package main

import "sort"

func maxOperations(nums []int, k int) int {

	sort.Ints(nums)
	count := 0
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] == k {
			i++
			j--
			count++
		} else if nums[i]+nums[j] > k {
			j--
		} else {
			i++
		}
	}
	return count
}
