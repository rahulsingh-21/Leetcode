package main

func countComponents(n int, edges [][]int) int {
	adj := map[int][]int{}
	visited := make([]bool, n)
	count := 0

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

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
