package main

func characterReplacement(s string, k int) int {

	l, r, ans := 0, 0, 0
	maxf := 0
	m := map[byte]int{}

	for r < len(s) {
		m[s[r]]++

		maxf = max(maxf, m[s[r]])

		if r-l+1-maxf > k {
			m[s[l]]--
			l++
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}

		r++
	}

	return ans
}
