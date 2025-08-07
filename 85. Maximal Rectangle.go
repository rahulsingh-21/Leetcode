package main

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix[0])
	heights := make([]int, n)
	maxArea := 0

	for i := range matrix {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		maxArea = max(maxArea, getMaxArea(heights))
	}

	return maxArea
}

func getMaxArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	heights = append(heights, 0)

	for i, h := range heights {
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}
	return maxArea
}
