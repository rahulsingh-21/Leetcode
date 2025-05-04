package main

import "strings"

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) > 0 {
			return len(arr[i])
		}
	}
	return 0
}
