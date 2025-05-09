package main

func isIsomorphic(s string, t string) bool {
	m1, m2 := map[byte]int{}, map[byte]int{}
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			// fmt.Println(m1,m2)
			return false
		} else {
			m1[s[i]] = i + 1
			m2[t[i]] = i + 1
		}

	}
	return true

}
