package main

func countCompleteComponents(n int, edges [][]int) int {
	m := make(map[int][]int)
	count := 0
	visited := make([]bool, n)

	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
		m[e[1]] = append(m[e[1]], e[0])
	}
	node, edge := 0, 0
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		node++
		for _, val := range m[i] {
			edge++
			if visited[val] {
				continue
			}
			dfs(val)
		}
	}
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		node, edge = 0, 0
		dfs(i)
		if edge%node == 0 && edge/node == node-1 {
			count++
		}
	}
	return count
}
