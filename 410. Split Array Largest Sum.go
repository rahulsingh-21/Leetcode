package main

func splitArray(nums []int, k int) int {
	sum, maxNum := 0, 0

	for _, num := range nums {
		sum += num
		maxNum = max(num, maxNum)
	}

	l, r := maxNum, sum

	for l < r {
		mid := (l + r) / 2
		if canSplit(nums, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func canSplit(nums []int, k int, maxAllowed int) bool {
	count := 1
	currSum := 0
	for _, num := range nums {
		if num+currSum > maxAllowed {
			currSum = num
			count++
			if count > k {
				return false
			}
		} else {
			currSum += num
		}
	}
	return true
}
