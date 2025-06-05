package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	n := len(word1)
	m := len(word2)

	if n != m {
		return false
	}

	m1, m2 := map[rune]int{}, map[rune]int{}

	for i, char := range word1 {
		m1[char]++
		m2[rune(word2[i])]++
	}

	for i, char := range word2 {
		if _, ok := m1[char]; !ok {
			return false
		}
		if _, ok := m2[rune(word1[i])]; !ok {
			return false
		}
	}

	freq1 := make([]int, len(m1))
	freq2 := make([]int, len(m2))

	for _, val := range m1 {
		freq1 = append(freq1, val)
	}

	for _, val := range m2 {
		freq2 = append(freq2, val)
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true

}
