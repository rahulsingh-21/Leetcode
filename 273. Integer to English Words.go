package main

import "strings"

func numberToWords(num int) string {

	if num == 0 {
		return "Zero"
	}
	i := 0
	result := ""

	for num > 0 {
		if num%1000 != 0 {
			result = helper(num%1000) + thousands[i] + " " + result
		}
		i++
		num /= 1000
	}
	return strings.TrimSpace(result)
}

func helper(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 {
		return below20[num] + " "
	} else if num < 100 {
		return tens[num/10] + " " + helper(num%10)
	} else {
		return helper(num/100) + "Hundred " + helper(num%100)
	}
}

var below20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

var thousands = []string{"", "Thousand", "Million", "Billion"}
