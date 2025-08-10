package main

func shipWithinDays(weights []int, days int) int {
	sum, maxNum := 0, 0
	for _, num := range weights {
		sum += num
		maxNum = max(num, maxNum)
	}

	l, r := maxNum, sum

	for l < r {
		mid := (l + r) / 2
		if canSplitII(weights, days, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func canSplitII(weights []int, days int, maxAllowed int) bool {
	count := 1
	curr := 0

	for _, weight := range weights {
		if curr+weight > maxAllowed {
			curr = weight
			count++
			if count > days {
				return false
			}
		} else {
			curr += weight
		}
	}
	return true
}
