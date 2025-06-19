package main

import "math"

func reverseII(x int) int {
	isN := 1
	if x < 0 {
		isN = -1
		x = -x
	}
	ans := 0
	for x > 0 {
		rem := x % 10
		ans = ans*10 + rem
		x /= 10
	}

	if ans > math.MaxInt32 {
		return 0
	}
	return ans * isN
}
