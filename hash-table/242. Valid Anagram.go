package hashtable

func isAnagram(s string, t string) bool {
	freq := make([]int, 26)
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		freq[t[i]-'a']--
		if freq[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
