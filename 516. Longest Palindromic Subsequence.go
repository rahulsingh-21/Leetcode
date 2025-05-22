package main

func longestPalindromeSubseq(s string) int {
	r := revI(s)
	n := len(s)

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[j-1] == r[i-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		prev, curr = curr, prev
	}
	return prev[n]

}

func revI(str string) string {
	i, j := 0, len(str)-1
	s := []rune(str)
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
