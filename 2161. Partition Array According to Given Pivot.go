package main

func pivotArray(nums []int, pivot int) []int {
	var less []int
	var equal []int
	var more []int

	for _, val := range nums {
		if val > pivot {
			more = append(more, val)
		} else if val < pivot {
			less = append(less, val)
		} else {
			equal = append(equal, val)
		}
	}
	less = append(append(less, equal...), more...)
	//less = append(less,more...)
	return less
}
