package main

import "unicode"

func isPalindromeII(s string) bool {

	str := ""
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			str += string(char)
		} else if char >= 'A' && char <= 'Z' {
			temp := unicode.ToLower(char)
			str += string(temp)
		}
	}

	i, j := 0, len(str)-1

	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
