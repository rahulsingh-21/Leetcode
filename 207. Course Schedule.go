package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := map[int][]int{}
	ind := make([]int, numCourses)

	for _, edge := range prerequisites {
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		ind[edge[0]]++
	}

	count := numCourses
	for count > 0 {
		idx := getZero(ind)
		if idx == -1 {
			return false
		}

		count--

		for _, edge := range adj[idx] {
			ind[edge]--
		}
	}

	return true

}

func getZero(n []int) int {
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			n[i] = -1
			return i
		}
	}
	return -1
}
