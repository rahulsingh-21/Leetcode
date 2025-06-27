package main

func canJump(nums []int) bool {
	maxi := 0
	for i, num := range nums {
		if maxi < i {
			return false
		}
		maxi = max(maxi, num+i)
	}
	return true
}
