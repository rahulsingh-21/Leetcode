package main

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	var res []int

	for l <= r {
		if nums[l]*nums[l] < nums[r]*nums[r] {
			res = append(res, nums[r]*nums[r])
			r--
		} else {
			res = append(res, nums[l]*nums[l])
			l++
		}
	}

	l, r = 0, len(nums)-1

	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
