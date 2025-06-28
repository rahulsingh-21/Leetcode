package main

func partitionLabels(s string) []int {
	m := map[rune]int{}

	for i, char := range s {
		m[char] = i
	}

	var ans []int

	start, end := 0, 0

	for i, char := range s {
		end = max(end, m[char])
		if i == end {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}
	return ans

}
