package main

import "math"

func increasingTriplet(nums []int) bool {
	//n := len(nums)

	f, s := math.MaxInt, math.MaxInt

	for _, n := range nums {
		if n <= f {
			f = n
		} else if n <= s {
			s = n
		} else if n > s && n > f {
			return true
		}
	}
	return false
}
