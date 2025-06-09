package main

func trap(height []int) int {

	n := len(height)
	left, right := make([]int, n), make([]int, n)

	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += min(left[i], right[i]) - height[i]
	}

	return ans
}
