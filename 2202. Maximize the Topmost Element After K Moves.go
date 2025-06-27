package main

import "slices"

func maximumTop(nums []int, k int) int {
	maxInBag := 0
	n := len(nums)

	if k == 0 {
		return nums[0]
	}

	if n == 1 && k%2 == 1 {
		return -1
	}

	if k > n {
		return slices.Max(nums)
	}

	for i := 1; i < k; i++ {
		maxInBag = max(maxInBag, nums[i-1])
	}

	if k == n {
		return maxInBag
	}

	return max(maxInBag, nums[k])
}
