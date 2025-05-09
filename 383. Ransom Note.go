package main

func canConstruct(ransomNote string, magazine string) bool {
	m := map[string]int{}

	if len(magazine) < len(ransomNote) {
		return false
	}

	for _, val := range magazine {
		m[string(val)]++
	}
	for _, val := range ransomNote {
		m[string(val)]--
		if m[string(val)] < 0 {
			return false
		}
	}
	return true

}
