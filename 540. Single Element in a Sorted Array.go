package main

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if mid%2 != 0 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			r = mid
		}
	}
	return nums[l]
}
