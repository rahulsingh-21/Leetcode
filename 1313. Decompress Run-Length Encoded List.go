package main

func decompressRLElist(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums)-1; i += 2 {
		count := nums[i]
		key := nums[i+1]
		for j := 0; j < count; j++ {
			ans = append(ans, key)
		}
	}
	return ans
}
