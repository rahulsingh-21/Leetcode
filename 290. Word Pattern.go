package main

import "strings"

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	m1, m2 := map[byte]int{}, map[string]int{}

	if len(pattern) != len(arr) {
		return false
	}

	for i := 0; i < len(arr); i++ {
		if m1[pattern[i]] != m2[arr[i]] {
			return false
		} else {
			m1[pattern[i]] = i + 1
			m2[arr[i]] = i + 1
		}
	}
	return true
}
