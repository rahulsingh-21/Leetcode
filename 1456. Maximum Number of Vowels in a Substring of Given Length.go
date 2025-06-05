package main

func maxVowels(s string, k int) int {
	count, ans := 0, 0
	l, r := 0, 0
	m := []byte{}
	for r < len(s) {
		m = append(m, s[r])
		if isVowel(s[r]) {
			count++
		}
		if len(m) > k {
			if isVowel(s[l]) {
				count--
			}
			m = m[1:]
			l++
		}
		r++
		ans = max(count, ans)

	}
	return ans
}

func isVowel(t byte) bool {
	s := rune(t)

	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}
	return false
}
