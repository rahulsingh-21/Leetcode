package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	prevEnd := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > prevEnd {
			count++
			prevEnd = points[i][1]
		}
	}
	return count
}
