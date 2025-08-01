package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	var backtrack func([]int, int, int)
	backtrack = func(curr []int, total int, start int) {
		if total == target {
			comb := append([]int{}, curr...)
			result = append(result, comb)
			return
		}

		if total > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(curr, total+candidates[i], i)
			curr = curr[:len(curr)-1]
		}
	}

	backtrack([]int{}, 0, 0)
	return result

}
