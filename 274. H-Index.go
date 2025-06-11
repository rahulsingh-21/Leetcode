package main

import "sort"

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	count := 0

	for i, n := range citations {
		if n >= i+1 {
			count++
		} else {
			return count
		}
	}
	return count
}
