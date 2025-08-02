package main

func permute(nums []int) [][]int {

	var result [][]int

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			perm := append([]int{}, nums...)
			result = append(result, perm)
			return
		}

		for i := index; i < len(nums); i++ {
			nums[index], nums[i] = nums[i], nums[index]
			backtrack(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}

	}
	backtrack(0)
	return result
}
