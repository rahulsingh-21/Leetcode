package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for _, num := range nums {
		sum += num
	}
	target := sum / k
	if nums[0] > target {
		return false
	}

	bucket := make([]int, k)
	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == len(nums) {
			return true
		}
		num := nums[idx]
		for i := 0; i < k; i++ {
			if bucket[i]+num <= target {
				bucket[i] += num
				if backtrack(idx + 1) {
					return true
				}
				bucket[i] -= num
			}

			if bucket[i] == 0 {
				break
			}
		}
		return false
	}

	return backtrack(0)
}
