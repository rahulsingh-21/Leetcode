package main

func isSubsequence(s string, t string) bool {
	i := 0
	n := len(s)
	if n == 0 {
		return true
	}
	for _, str := range t {
		if str == rune(s[i]) {
			i++
		}
		if i == n {
			return true
		}
	}

	return false
}
