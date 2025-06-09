package main

func lengthOfLongestSubstring(s string) int {
	ans := 0
	l, r := 0, 0

	m := map[byte]int{}

	for r < len(s) {
		m[s[r]]++
		if m[s[r]] > 1 {
			for m[s[r]] > 1 {
				m[s[l]]--
				l++
			}
			//delete(m, s[r])

		}
		if r-l+1 > ans {
			ans = r - l + 1
		}
		r++
	}

	return ans
}
