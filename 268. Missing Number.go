package main

func missingNumber(nums []int) int {
	freq := make([]int, len(nums)+1)

	for _, val := range nums {
		freq[val]++
	}
	for i, val := range freq {
		if val == 0 {
			return i
		}
	}
	return -1
}
