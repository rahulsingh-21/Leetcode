package main

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	i, j := 0, m-1

	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}
