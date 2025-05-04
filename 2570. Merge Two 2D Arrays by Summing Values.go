package main

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		m[nums1[i][0]] += nums1[i][1]
	}
	for i := 0; i < len(nums2); i++ {
		m[nums2[i][0]] += nums2[i][1]
	}

	var n []int
	for key, _ := range m {
		n = append(n, key)
	}
	sort.Ints(n)
	var ans [][]int
	for i := 0; i < len(n); i++ {
		var temp []int
		temp = append(temp, n[i], m[n[i]])
		ans = append(ans, temp)
	}
	return ans
}
