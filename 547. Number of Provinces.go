package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	adj := map[int][]int{}
	visited := make([]bool, n+1)
	count := 0

	for i, edge := range isConnected {
		for j, eds := range edge {
			if i != j && eds == 1 {
				fmt.Println(i, j)
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	fmt.Println(adj)
	var dfs func(n int)

	dfs = func(n int) {
		visited[n] = true

		for _, edge := range adj[n] {
			if !visited[edge] {
				dfs(edge)
			}
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		dfs(i)
		count++
	}
	return count
}
