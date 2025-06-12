package main

func checkValidString(s string) bool {
	openCount := 0
	closeCount := 0
	n := len(s) - 1
	for i := 0; i <= n; i++ {
		if s[i] == '(' || s[i] == '*' {
			openCount++
		} else {
			openCount--
		}

		if s[n-i] == ')' || s[n-i] == '*' {
			closeCount++
		} else {
			closeCount--
		}

		if openCount < 0 || closeCount < 0 {
			return false
		}

	}

	return true
}
