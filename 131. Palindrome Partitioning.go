package main

func partition(s string) [][]string {

	var result [][]string
	var backtrack func(start int, curr []string)
	backtrack = func(start int, curr []string) {
		if start >= len(s) {
			temp := append([]string{}, curr...)
			result = append(result, temp)
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if checkV(str) {
				curr = append(curr, str)
				backtrack(i+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}
	backtrack(0, []string{})
	return result
}

func checkV(str string) bool {
	l, r := 0, len(str)-1

	for l < r {
		if str[l] == str[r] {
			r--
			l++
		} else {
			return false
		}
	}
	return true
}
