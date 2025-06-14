package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {

	n := make([]int, len(nums))
	copy(n, nums)
	sort.Ints(n)

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[n[i]]; !ok {
			m[n[i]] = i
		}
	}

	ans := make([]int, len(nums))

	fmt.Println(nums, m, n, ans)
	for i := 0; i < len(nums); i++ {
		ans[i] = m[nums[i]]
	}

	return ans
}
