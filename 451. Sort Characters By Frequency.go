package main

import "sort"

func frequencySort(s string) string {
	freq := make([]int, 255)
	for _, n := range s {
		freq[n]++
	}

	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		if freq[str[i]] == freq[str[j]] {
			return str[i] > str[j]
		}
		return freq[str[i]] > freq[str[j]]
	})
	return string(str)
}
