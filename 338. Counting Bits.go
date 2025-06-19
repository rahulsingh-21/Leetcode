package main

func countBits(n int) []int {
	var ans []int

	for i := 0; i <= n; i++ {
		ans = append(ans, count(i))
	}
	return ans
}

func count(n int) int {
	ans := 0
	for n > 0 {
		ans += (n & 1)
		n = n >> 1
	}
	return ans
}
