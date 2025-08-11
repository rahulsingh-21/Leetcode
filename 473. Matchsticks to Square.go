package main

import "sort"

func makesquare(matchsticks []int) bool {
	sum := 0

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	for _, num := range matchsticks {
		sum += num
	}

	target := sum / 4
	if target < matchsticks[0] {
		return false
	}

	bucket := make([]int, 4)

	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		num := matchsticks[idx]
		for i := 0; i < 4; i++ {
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
