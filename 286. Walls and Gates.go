package main

import "math"

func wallsAndGates(rooms [][]int) {
	empty := math.MaxInt32
	gate := 0
	//wall := -1
	n := len(rooms)
	m := len(rooms[0])

	var queue [][]int

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rooms[i][j] == gate {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			x, y := room[0]+dir[0], room[1]+dir[1]
			if x >= 0 && x < n && y >= 0 && y < m && rooms[x][y] == empty {
				rooms[x][y] = rooms[room[0]][room[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
}
