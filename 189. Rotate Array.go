package twopointers

func rotate(nums []int, k int) {
	k = k % len(nums)
	rev(nums)
	rev(nums[:k])
	rev(nums[k:])
}

func rev(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
