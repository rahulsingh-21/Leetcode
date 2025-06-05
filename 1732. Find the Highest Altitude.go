package main

func largestAltitude(gain []int) int {
	ans, sum := 0, 0
	for _, v := range gain {
		sum += v
		ans = max(sum, ans)
	}
	return ans
}
