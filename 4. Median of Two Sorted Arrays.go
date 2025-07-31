package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	merged := make([]int, 0, m+n)
	i, j := 0, 0

	for i < n && j < m {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	for i < n {
		merged = append(merged, nums1[i])
		i++
	}

	for j < m {
		merged = append(merged, nums2[j])
		j++
	}

	mid := len(merged) / 2
	if len(merged)%2 == 0 {
		return (float64(merged[mid-1]) + float64(merged[mid])) / 2.0
	} else {
		return float64(merged[mid])
	}
}
