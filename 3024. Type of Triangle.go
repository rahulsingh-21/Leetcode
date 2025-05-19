package main

import "sort"

func triangleType(nums []int) string {
	m := map[int]struct{}{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}

	if len(m) == 1 {
		return "equilateral"
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	}
	if len(m) == 2 {
		return "isosceles"
	}

	return "scalene"
}
