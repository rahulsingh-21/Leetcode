package main

func totalNQueens(n int) int {

	count := 0
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
			count++
			return
		}

		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[col+row] {
				continue
			}

			board[row][col] = 'Q'
			cols[col], diag1[row-col], diag2[col+row] = true, true, true
			backtrack(row + 1)
			board[row][col] = '.'
			cols[col], diag1[row-col], diag2[col+row] = false, false, false
		}
	}
	backtrack(0)
	return count
}
