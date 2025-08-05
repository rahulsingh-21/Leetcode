package main

func searchRange(nums []int, target int) []int {
	return []int{left(nums, target), rightI(nums, target)}
}

func left(nums []int, target int) int {
	l, r := 0, len(nums)-1
	index := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			index = mid
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}

func rightI(nums []int, target int) int {
	l, r := 0, len(nums)-1
	index := -1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			index = mid
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return index
}
