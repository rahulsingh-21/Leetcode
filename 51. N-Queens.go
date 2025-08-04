package main

func solveNQueens(n int) [][]string {

	var result [][]string
	board := make([][]byte, n)

	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	cols := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	var backtrack func(int)
	backtrack = func(row int) {

		if row == n {
			var sol []string
			for _, piece := range board {
				sol = append(sol, string(piece))
			}
			result = append(result, sol)
			return
		}

		for col := 0; col < n; col++ {

			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}

			board[row][col] = 'Q'
			cols[col], diag1[row-col], diag2[row+col] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[row-col], diag2[row+col] = false, false, false
		}
	}

	backtrack(0)
	return result
}
