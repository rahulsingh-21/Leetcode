package main

import "strconv"

func equalPairs(grid [][]int) int {
	rowMap := map[string]int{}
	colMap := map[string]int{}

	n := len(grid)

	for i := 0; i < n; i++ {
		col := ""
		row := ""
		for j := 0; j < n; j++ {
			row += strconv.Itoa(grid[i][j]) + ","
			col += strconv.Itoa(grid[j][i]) + ","
		}
		rowMap[row]++
		colMap[col]++
	}

	count := 0
	for key, value := range rowMap {
		count += value * colMap[key]
	}

	return count
}
