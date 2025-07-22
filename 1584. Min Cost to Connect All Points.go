package main

import "sort"

func minCostConnectPoints(points [][]int) int {

	var edges []Edge
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1 := points[i][0]
			y1 := points[i][1]
			x2 := points[j][0]
			y2 := points[j][1]
			dis := abs(x1-x2) + abs(y1-y2)
			edges = append(edges, Edge{i, j, dis})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	parent := make([]int, n)
	rank := make([]int, n)

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

	var union func(int, int) bool
	union = func(a, b int) bool {
		u, v := find(a), find(b)
		if u == v {
			return false
		}
		if rank[u] < rank[v] {
			parent[u] = v
		} else if rank[v] < rank[u] {
			parent[v] = u
		} else {
			parent[u] = v
			rank[v]++
		}
		return true
	}

	minCost := 0
	for _, edge := range edges {
		if union(edge.x, edge.y) {
			minCost += edge.cost
		}
	}

	return minCost
}

type Edge struct {
	x, y, cost int
}
