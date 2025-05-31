package main

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	q := [][]int{}
	q = append(q, []int{0, 0, 2})

	for len(q) > 0 {
		for _, cell := range q {
			for _, dir := range direction {
				x := cell[0] + dir[0]
				y := cell[1] + dir[1]
				path := cell[2]

				if x == n-1 && y == n-1 {
					return path
				}

				if x < n && x >= 0 && y >= 0 && y < n && grid[x][y] == 0 {
					q = append(q, []int{x, y, path + 1})
					grid[x][y] = 1

				}
			}
			q = q[1:]
		}
	}
	return -1
}

var direction = [8][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
