package main

func orangesRotting(grid [][]int) int {

	fresh := 0
	q := [][]int{}
	n := len(grid)
	m := len(grid[0])
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	if fresh == 0 {
		return 0
	}

	for len(q) > 0 {
		for _, node := range q {
			for _, dir := range directions {
				x, y := node[0]+dir[0], node[1]+dir[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 1 {
					q = append(q, []int{x, y})
					grid[x][y] = 2
					fresh--
				}
			}
			q = q[1:]
		}
		count++
		if fresh == 0 {
			return count
		}
	}
	return -1

}
