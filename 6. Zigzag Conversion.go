package main

import "strings"

func convert(s string, numRows int) string {
	rowIndex := 0
	str := make([]string, numRows)
	dir := 1

	if numRows == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {

		str[rowIndex] = str[rowIndex] + string(s[i])
		rowIndex += dir

		if rowIndex == numRows-1 || rowIndex == 0 {
			dir = -dir
		}
	}

	return strings.Join(str, "")
}
