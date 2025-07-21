package main

func sortColors(nums []int) {
	p0, p2, curr := 0, len(nums)-1, 0

	for curr <= p2 {
		if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr]
			curr++
			p0++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		} else {
			curr++
		}
	}
}
