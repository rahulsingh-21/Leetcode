package main

func containsDuplicate(nums []int) bool {

	set := map[int]struct{}{}

	for _, n := range nums {
		set[n] = struct{}{}
	}

	return !(len(set) == len(nums))
}
