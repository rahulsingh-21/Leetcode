package main

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	n := len(text1)
	m := len(text2)

	curr := make([]int, n+1)
	prev := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text2[i-1] == text1[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(curr[j-1], prev[j])
			}
		}
		curr, prev = prev, curr
	}

	return prev[n]
}
