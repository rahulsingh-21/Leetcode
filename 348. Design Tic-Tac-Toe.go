package main

type TicTacToe struct {
	rows         [][]int
	columns      [][]int
	diagonal     []int
	antidiagonal []int
	size         int
}

func ConstructorXII(n int) TicTacToe {
	rows := make([][]int, 2)
	columns := make([][]int, 2)

	for i := 0; i < 2; i++ {
		rows[i] = make([]int, n)
		columns[i] = make([]int, n)
	}
	diagonal := make([]int, 2)
	antidiagonal := make([]int, 2)
	return TicTacToe{
		rows, columns, diagonal, antidiagonal, n,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	playerIdx := player - 1
	this.rows[playerIdx][row]++
	this.columns[playerIdx][col]++

	if row == col {
		this.diagonal[playerIdx]++
	}

	if row+col == this.size-1 {
		this.antidiagonal[playerIdx]++
	}

	if this.rows[playerIdx][row] == this.size ||
		this.columns[playerIdx][col] == this.size ||
		this.diagonal[playerIdx] == this.size ||
		this.antidiagonal[playerIdx] == this.size {
		return player
	}

	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
