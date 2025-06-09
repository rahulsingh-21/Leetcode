package main

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			twoSumIII(nums, i, &ans)
		}
	}
	return ans
}

func twoSumIII(nums []int, i int, ans *[][]int) {
	l, m := i+1, len(nums)-1

	for l < m {
		if nums[i]+nums[l]+nums[m] == 0 {
			*ans = append(*ans, []int{nums[i], nums[l], nums[m]})
			l++
			m--
			for l < m && nums[l] == nums[l-1] {
				l++
			}
		} else if nums[i]+nums[l]+nums[m] > 0 {
			m--
		} else {
			l++
		}
	}
}
