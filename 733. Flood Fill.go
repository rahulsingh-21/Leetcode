package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var col [][]int
	col = append(col, []int{sr, sc})
	n, m := len(image), len(image[0])
	og := image[sr][sc]
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if image[i][j] == image[sr][sc] {
				count++
			}
		}
	}
	if color == image[sr][sc] {
		return image
	}
	if count == 0 {
		return image
	}
	image[sr][sc] = color

	for len(col) > 0 {
		for _, c := range col {
			for _, dir := range directions {
				i, j := c[0]+dir[0], c[1]+dir[1]
				if i >= 0 && i < n && j >= 0 && j < m && image[i][j] == og {
					image[i][j] = color
					fmt.Println(c, dir)
					col = append(col, []int{i, j})
					count--
				}
			}
			col = col[1:]
		}
	}
	return image
}

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
