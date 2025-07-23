package main

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for j := 0; j < len(rooms[i]); j++ {
			if !visited[rooms[i][j]] {
				dfs(rooms[i][j])
			}
		}
	}

	dfs(0)

	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}

	return true

}
