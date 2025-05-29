package main

func updateMatrix(mat [][]int) [][]int {
	n := len(mat)
	m := len(mat[0])
	var bfs [][]int
	//bfs = append(bfs,[]int{0,0})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 0 {
				bfs = append(bfs, []int{i, j})
			} else {
				mat[i][j] = n*m + 1
			}
		}
	}

	for len(bfs) > 0 {
		for _, b := range bfs {
			//count := 0
			for _, d := range directions {
				x, y := b[0], b[1]
				i, j := b[0]+d[0], b[1]+d[1]
				for i >= 0 && i < n && j >= 0 && j < m && mat[i][j] > mat[x][y]+1 {
					mat[i][j] = mat[x][y] + 1
					bfs = append(bfs, []int{i, j})
				}
			}
			bfs = bfs[1:]
		}
	}
	return mat
}
