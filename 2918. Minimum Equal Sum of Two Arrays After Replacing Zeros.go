package main

func minSum(nums1 []int, nums2 []int) int64 {
	var zero1, zero2, sum1, sum2 int64

	for _, val := range nums1 {
		if val == 0 {
			sum1 += 1
			zero1 += 1
		}
		sum1 += int64(val)
	}

	for _, val := range nums2 {
		if val == 0 {
			sum2 += 1
			zero2 += 1
		}
		sum2 += int64(val)
	}

	if zero2 == 0 && sum1 > sum2 {
		return -1
	}
	if zero1 == 0 && sum2 > sum1 {
		return -1
	}
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}
