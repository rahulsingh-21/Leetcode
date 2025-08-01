package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)

	var backtrack func(int, int, []int)
	backtrack = func(start, total int, curr []int) {
		if total == target {
			comb := append([]int{}, curr...)
			result = append(result, comb)
			return
		}

		if total > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			curr = append(curr, candidates[i])
			backtrack(i+1, total+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack(0, 0, []int{})

	return result
}
