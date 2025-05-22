package main

func minInsertions(s string) int {
	r := revI(s)
	n := len(s)

	curr := make([]int, n+1)
	prev := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if r[j-1] == s[i-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(curr[j-1], prev[j])
			}
		}
		curr, prev = prev, curr
	}
	return n - prev[n]
}
