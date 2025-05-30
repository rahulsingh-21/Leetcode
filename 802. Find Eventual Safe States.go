package main

import "sort"

func eventualSafeNodes(graph [][]int) []int {
	ans := []int{}
	indegree := make([]int, len(graph))
	ad := make(map[int][]int, len(graph))

	for i, val := range graph {
		for _, e := range val {
			ad[e] = append(ad[e], i)
			indegree[i]++
		}
	}

	count := len(graph)

	for count > 0 {
		i := checkZero(indegree)
		if i == -1 {
			sort.Slice(ans, func(i, j int) bool {
				return ans[i] < ans[j]
			})
			return ans
		}
		count--
		ans = append(ans, i)

		for _, val := range ad[i] {
			indegree[val]--

		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func checkZero(n []int) int {
	for i, val := range n {
		if val == 0 {
			n[i] = -1
			return i
		}
	}
	return -1
}
