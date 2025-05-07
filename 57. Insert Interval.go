package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var ans [][]int

	for _, val := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < val[0] {
			ans = append(ans, val)
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], val[1])
		}
	}

	return ans
}
