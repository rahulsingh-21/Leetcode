package main

func pacificAtlantic(heights [][]int) [][]int {
	n := len(heights)
	m := len(heights[0])

	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)

	for i := range pacific {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)
	}

	var dfs func(int, int, [][]bool)
	dfs = func(i, j int, visited [][]bool) {
		visited[i][j] = true
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if x < 0 || x >= n || y < 0 || y >= m || visited[x][y] || heights[x][y] < heights[i][j] {
				continue
			}
			dfs(x, y, visited)
		}
	}

	for i := 0; i < n; i++ {
		dfs(i, 0, pacific)    //pacific left edge
		dfs(i, m-1, atlantic) // atlantic right edge
	}

	for j := 0; j < m; j++ {
		dfs(0, j, pacific)    // pacific top edge
		dfs(n-1, j, atlantic) // atlantic bottom edge
	}

	var result [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
