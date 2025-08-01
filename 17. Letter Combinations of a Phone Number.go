package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var result []string
	var backtrack func(int, string)
	backtrack = func(index int, curr string) {
		if index == len(digits) {
			result = append(result, curr)
			return
		}

		letters := phone[digits[index]]
		for _, letter := range letters {
			backtrack(index+1, curr+string(letter))
		}
	}

	backtrack(0, "")
	return result
}
