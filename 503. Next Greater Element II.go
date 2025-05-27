package main

func nextGreaterElements(nums []int) []int {
	//var ans []int
	nt := make([]int, len(nums))
	nums2 := nums
	nums2 = append(nums2, nums...)
	//fmt.Println(nums2)
	//return []int{}

	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums[i] {
				nt[i] = nums2[j]
				flag = true
				break
			}
		}
		if !flag {
			nt[i] = -1
		}
	}
	return nt
}
