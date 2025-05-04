package main

func largestOddNumber(num string) string {
	// ans := ""
	index := len(num) - 1
	for index = len(num) - 1; index >= 0 && num[index]%2 == 0; {
		index--
	}

	return num[:index+1]
}
