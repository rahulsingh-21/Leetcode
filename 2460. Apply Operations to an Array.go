package array

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	ans := make([]int, len(nums))
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ans[k] = nums[i]
			k++
		}
	}
	return ans
}
