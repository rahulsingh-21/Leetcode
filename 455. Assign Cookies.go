package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	cg, cs := 0, 0
	for cg < len(g) && cs < len(s) {
		if s[cs] >= g[cg] {
			cg++
		}
		cs++
	}
	return cg
}
