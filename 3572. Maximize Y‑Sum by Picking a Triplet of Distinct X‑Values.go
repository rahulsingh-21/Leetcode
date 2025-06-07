package main

import (
	"slices"
	"sort"
)

func maxSumDistinctTriplet(x []int, y []int) int {

	n := len(x)
	var pair []Pair
	for i := 0; i < n; i++ {
		pair = append(pair, Pair{x[i], y[i]})
	}

	sort.Slice(pair, func(i, j int) bool {
		return pair[i].b > pair[j].b
	})

	count := 0
	ans := 0
	// temp1 := 0
	//temp2 := 0
	var temp []int
	for i := 0; i < len(pair); i++ {
		if count == 3 {
			return ans
		}
		// fmt.Println(temp1, temp)
		if i == 0 {
			ans += pair[i].b
			//temp1 = pair[i].b
			temp = append(temp, pair[i].a)
			count++
		}
		if !slices.Contains(temp, pair[i].a) {
			ans += pair[i].b
			//temp1 = pair[i].b
			temp = append(temp, pair[i].a)
			count++
		}
	}
	if count == 3 {
		return ans
	}

	return -1

}

type Pair struct {
	a, b int
}
