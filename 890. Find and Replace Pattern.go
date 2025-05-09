package main

func findAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		m1 := map[byte]int{}
		m2 := map[byte]int{}
		flag := true
		if len(word) != len(pattern) {
			continue
		}
		for i := 0; i < len(word); i++ {
			if m1[word[i]] != m2[pattern[i]] {
				flag = false
			} else {
				m1[word[i]] = i + 1
				m2[pattern[i]] = i + 1
			}
		}
		if flag {
			ans = append(ans, word)
		}
	}
	return ans
}
