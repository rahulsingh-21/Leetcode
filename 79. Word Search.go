package main

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	var backtrack func(int, int, int) bool
	backtrack = func(i, j, index int) bool {
		if i < 0 || j < 0 || i > n-1 || j > m-1 || board[i][j] == '0' {
			return false
		}

		if board[i][j] != word[index] {
			return false
		}

		if index == len(word)-1 {
			return board[i][j] == word[index]
		}

		og := board[i][j]
		board[i][j] = '0'

		if backtrack(i+1, j, index+1) ||
			backtrack(i-1, j, index+1) ||
			backtrack(i, j+1, index+1) ||
			backtrack(i, j-1, index+1) {
			return true
		}

		board[i][j] = og

		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}
