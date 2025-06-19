package main

func hammingWeight(n int) int {
	ans := 0
	for n > 0 {
		ans += int(n & 1)
		n = n >> 1
	}
	return ans
}
