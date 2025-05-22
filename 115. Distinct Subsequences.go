package main

func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)

	if m > n {
		return 0
	}

	prev := make([]int, m+1)
	curr := make([]int, m+1)

	prev[0] = 1

	for i := 1; i <= n; i++ {
		curr[0] = 1
		for j := 1; j <= m; j++ {
			if t[j-1] == s[i-1] {
				curr[j] = prev[j-1] + prev[j]
			} else {
				curr[j] = prev[j]
			}
		}
		prev, curr = curr, prev
	}

	return prev[m]
}
