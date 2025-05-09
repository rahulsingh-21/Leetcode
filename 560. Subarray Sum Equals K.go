package main

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	sum, count := 0, 0
	for _, val := range nums {
		sum += val
		count += m[sum-k]

		m[sum]++
	}
	return count
}
