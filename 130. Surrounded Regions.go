package main

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])
	q := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' && (i == 0 || j == 0 || i == n-1 || j == m-1) {
				board[i][j] = 'R'
				q = append(q, []int{i, j})
			}
		}
	}

	for len(q) > 0 {
		for _, node := range q {
			for _, dir := range directions {
				i := node[0] + dir[0]
				j := node[1] + dir[1]
				if i >= 0 && i < n && j >= 0 && j < m && board[i][j] == 'O' {
					board[i][j] = 'R'
					q = append(q, []int{i, j})
				}
			}
			q = q[1:]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			}
		}
	}
	return
}
