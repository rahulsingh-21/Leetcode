package main

func setZeroes(matrix [][]int) {
	col := false

	r := len(matrix)
	c := len(matrix[0])

	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			col = true
		}
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < c; i++ {
			matrix[0][i] = 0
		}
	}

	if col {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}
}
