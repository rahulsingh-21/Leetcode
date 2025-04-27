package twopointers

func removeDuplicates(nums []int) int {
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}
