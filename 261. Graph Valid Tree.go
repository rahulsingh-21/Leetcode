package main

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		u, v := find(x), find(y)
		if u == v {
			return false
		}
		parent[u] = v
	}
	return true
}
