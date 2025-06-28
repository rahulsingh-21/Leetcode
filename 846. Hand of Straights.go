package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}
	m := map[int]int{}

	sort.Ints(hand)
	for _, g := range hand {
		m[g]++
	}

	for _, h := range hand {
		curr := h
		if m[h] == 0 {
			continue
		}
		n := 0
		for n < groupSize {
			if m[curr] == 0 {
				return false
			}
			m[curr]--
			curr++
			n++
		}
	}

	return true
}
