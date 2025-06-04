package main

import "strings"

func reverseWords(s string) string {
	strArr := strings.Fields(s)

	for i, j := 0, len(strArr)-1; i < j; {
		strArr[i], strArr[j] = strArr[j], strArr[i]
		i++
		j--
	}

	return strings.Join(strArr, " ")
}
