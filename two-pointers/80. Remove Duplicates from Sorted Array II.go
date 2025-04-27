package twopointers

func removeDuplicates2(nums []int) int {
	ans, c := 0, 0
	nums[ans] = nums[0]
	ans++
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			c = 0
		} else {
			c++
		}
		if c >= 2 {
			continue
		}
		nums[ans] = nums[i]
		ans++
	}
	return ans
}
