package main

func maxArea(height []int) int {

	ans := 0
	for i, j := 0, len(height)-1; i < j; {
		diff := min(height[i], height[j]) * (j - i)
		ans = max(ans, diff)
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return ans
}
