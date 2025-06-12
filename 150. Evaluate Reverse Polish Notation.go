package main

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int

	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {

			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			var temp int
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				temp = a + b
			case "-":
				temp = b - a
			case "*":
				temp = b * a
			case "/":
				temp = b / a
			}
			stack = append(stack, temp)

		} else {
			tokenInt, _ := strconv.Atoi(token)
			stack = append(stack, tokenInt)
		}
	}

	return stack[0]
}
