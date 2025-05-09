package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}

	for i, val := range nums {
		if index, ok := m[val]; ok {
			if i-index <= k {
				return true
			}
		}
		m[val] = i
	}
	return false
}
