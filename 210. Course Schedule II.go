package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := map[int][]int{}
	ind := make([]int, numCourses)

	for _, edge := range prerequisites {
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		ind[edge[0]]++
	}

	count := numCourses
	courses := []int{}
	for count > 0 {
		idx := getZero(ind)
		if idx == -1 {
			return []int{}
		}

		count--
		courses = append(courses, idx)

		for _, edge := range adj[idx] {
			ind[edge]--
		}
	}
	return courses
}
