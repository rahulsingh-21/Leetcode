package twopointers

func removeElement(nums []int, val int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}
