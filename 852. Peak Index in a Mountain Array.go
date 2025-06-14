package main

func peakIndexInMountainArray(arr []int) int {

	l, r := 0, len(arr)-1

	for l < r {
		mid := (r + l) / 2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
