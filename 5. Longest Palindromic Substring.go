package main

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		odd := check(i, i, s)
		if len(odd) > len(ans) {
			ans = odd
		}
		even := check(i, i+1, s)
		if len(even) > len(ans) {
			ans = even
		}
	}
	return ans
}

func check(l, r int, s string) string {

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
