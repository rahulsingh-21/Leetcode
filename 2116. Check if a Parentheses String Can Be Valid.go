package main

func canBeValid(s string, locked string) bool {

	open := 0
	close := 0

	n := len(s) - 1

	if len(s)%2 != 0 {
		return false
	}

	for i := 0; i <= n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			open++
		} else {
			open--
		}

		if s[n-i] == ')' || locked[n-i] == '0' {
			close++
		} else {
			close--
		}

		if open < 0 || close < 0 {
			return false
		}
	}

	return true
}
