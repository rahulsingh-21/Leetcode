package main

func rotateII(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		j := 0
		k := len(matrix[i]) - 1
		for j < k {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
			j++
			k--
		}
	}
}
