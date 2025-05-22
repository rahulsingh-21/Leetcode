package main

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	prev := make([]int, m+1)
	curr := make([]int, m+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word2[j-1] == word1[i-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		curr, prev = prev, curr
	}

	return n + m - 2*prev[m]
}
