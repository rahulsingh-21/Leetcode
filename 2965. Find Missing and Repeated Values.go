package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	m := make([]int, len(grid)*len(grid)+1)
	ans := make([]int, 2)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			m[grid[i][j]]++
		}
	}

	//fmt.Println(m)
	for i := 1; i < len(m); i++ {
		if m[i] == 2 {
			ans[0] = i
		} else if m[i] == 0 {
			ans[1] = i
		}
	}
	return ans
}
