package main

func numberOfSubstrings(s string) int {
	i := 0
	j := 0
	ans := 0
	m := make(map[byte]int)
	for i < len(s) {
		if val, ok := m[s[i]]; ok {
			m[s[i]] = val + 1
		} else {
			m[s[i]] = 1
		}

		for contains(m) {
			ans += len(s) - i
			m[s[j]] = m[s[j]] - 1
			j++
		}
		i++
	}
	return ans
}

func contains(m map[byte]int) bool {
	return m['a'] > 0 && m['b'] > 0 && m['c'] > 0
}
