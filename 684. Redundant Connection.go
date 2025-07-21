package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)

	parent := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var union func(int, int) bool
	union = func(a, b int) bool {
		u := find(a)
		v := find(b)
		if u == v {
			return false
		}
		parent[u] = v
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return []int{edge[0], edge[1]}
		}
	}
	return []int{}
}
