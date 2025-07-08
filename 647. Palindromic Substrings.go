package main

func countSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		count += countPalindrome(i, i, s) + countPalindrome(i, i+1, s)
	}
	return count
}

func countPalindrome(l, r int, s string) int {

	count := 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
