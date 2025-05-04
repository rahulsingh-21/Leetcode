package main

func majorityElement(nums []int) int {
	// max := 0
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
