package main

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	bucket := make([][]int, len(nums)+1)

	for _, num := range nums {
		m[num]++
	}

	for key, value := range m {
		bucket[value] = append(bucket[value], key)
	}

	var ans []int
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, n := range bucket[i] {
			if k > 0 {
				ans = append(ans, n)
				k--
			}
		}
	}
	return ans
}
