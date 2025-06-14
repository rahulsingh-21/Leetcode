package main

func searchMatrixII(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2
		val := matrix[mid/m][mid%m]

		if val == target {
			return true
		} else if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
