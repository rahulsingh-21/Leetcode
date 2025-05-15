package main

func getLargestOutlier(nums []int) int {
	freq := map[int]int{}
	sum, ans := 0, -10000
	for _, n := range nums {
		freq[n]++
		sum += n
	}

	for _, n := range nums {
		temp := sum - 2*n
		if freq[temp] > 0 {
			if temp != n || freq[temp] > 1 {
				ans = max(ans, temp)
			}
		}
	}
	return ans
}
