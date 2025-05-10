package main

import "fmt"

func maxSubArrayLen(nums []int, k int) int {
	m := map[int]int{0: -1}
	sum, ans := 0, 0
	for i, val := range nums {
		sum += val
		if idx, ok := m[sum-k]; ok {
			if i-idx > ans {
				fmt.Println(i, idx)
				ans = i - idx
			}
		}
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
		/// m[sum] = i
	}
	return ans
}
