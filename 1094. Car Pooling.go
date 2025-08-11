package main

func carPooling(trips [][]int, capacity int) bool {
	prefix := make([]int, 1001)

	for _, trip := range trips {
		passengers, from, to := trip[0], trip[1], trip[2]
		prefix[from] += passengers
		prefix[to] -= passengers
	}

	total := 0
	for i := 0; i < len(prefix); i++ {
		total += prefix[i]
		if total > capacity {
			return false
		}
	}
	return true
}

// 1 2 3 4 5 6 7
// 2 2 5 5 3 3 0
// 2 0 3 0 -2 0 -3
