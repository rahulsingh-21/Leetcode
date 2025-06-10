package main

func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	if n > m {
		return false
	}

	for i := 0; i < n; i++ {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}

	for i := 0; i < m-n; i++ {
		if match(freq1, freq2) {
			return true
		}
		freq2[s2[i]-'a']--
		freq2[s2[i+n]-'a']++
	}

	return match(freq1, freq2)
}

func match(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
