package main

func sortColors(nums []int) {
	c0, c2 := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		for (nums[i] == 0 || nums[i] == 2) && i >= c0 && c2 >= i {
			if nums[i] == 0 {
				nums[i], nums[c0] = nums[c0], nums[i]
				c0++
			} else if nums[i] == 2 {
				nums[i], nums[c2] = nums[c2], nums[i]
				c2--
			}
		}
	}
}
