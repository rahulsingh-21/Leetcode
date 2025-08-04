package main

func longestIncreasingPath(matrix [][]int) int {

	n := len(matrix)
	m := len(matrix[0])

	if n == 0 || m == 0 {
		return 0
	}

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, m)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if dp[i][j] != 0 {
			return dp[i][j]
		}

		maxLength := 1

		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && y >= 0 && x < n && y < m && matrix[i][j] < matrix[x][y] {
				length := 1 + dfs(x, y)
				maxLength = max(maxLength, length)
			}
		}

		dp[i][j] = maxLength
		return maxLength
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result = max(result, dfs(i, j))
		}
	}
	return result
}
