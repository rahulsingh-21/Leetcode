package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	count := 0

	sort.Slice(people, func(i, j int) bool {
		return people[i] < people[j]
	})

	i := 0
	j := len(people) - 1

	for i <= j {
		weight := people[i] + people[j]
		if weight > limit {
			count++
			j--
		} else {
			count++
			i++
			j--
		}
	}
	return count
}
