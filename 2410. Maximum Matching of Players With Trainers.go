package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Slice(players, func(i, j int) bool {
		return players[i] < players[j]
	})

	sort.Slice(trainers, func(i, j int) bool {
		return trainers[i] < trainers[j]
	})

	l, r := 0, 0

	for l < len(players) && r < len(trainers) {
		if trainers[r] >= players[l] {
			l++
		}
		r++
	}
	return l
}
