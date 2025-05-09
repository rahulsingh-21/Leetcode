package main

func twoSumII(nums []int, target int) []int {
	m := map[int]int{}

	for i, val := range nums {
		if idx, ok := m[target-val]; ok {
			return []int{idx, i}
		}

		m[val] = i
	}
	return nil
}
