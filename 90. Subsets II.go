package main

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	sort.Ints(nums)
	seen := map[string]bool{}
	for _, num := range nums {
		var set [][]int

		for _, curr := range result {
			newSet := append([]int{}, curr...)
			newSet = append(newSet, num)

			key := toKey(newSet)
			if !seen[key] {
				seen[key] = true
				set = append(set, newSet)
			}
		}
		result = append(result, set...)
	}
	return result
}

func toKey(nums []int) string {
	key := ""
	for i, n := range nums {
		if i > 0 {
			key += ","
		}
		key += strconv.Itoa(n)
	}
	return key
}
