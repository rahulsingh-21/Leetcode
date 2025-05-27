package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	var ans []int
	nt := make([]int, 10001)

	for i := 0; i < len(nums2); i++ {
		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				nt[nums2[i]] = nums2[j]
				break
			}
		}
	}
	//fmt.Println(nt)
	for i := 0; i < len(nums1); i++ {
		if nt[nums1[i]] == 0 {
			nt[nums1[i]] = -1
		}
		ans = append(ans, nt[nums1[i]])
	}
	return ans

}
