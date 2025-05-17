package main

func minPathSum(grid [][]int) int {
	for i := len(grid) - 2; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] += min(grid[i+1][j], grid[i+1][j+1])
		}
	}
	return grid[0][0]
}
