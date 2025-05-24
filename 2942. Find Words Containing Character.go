package main

func findWordsContaining(words []string, x byte) []int {

	var ans []int
	for i, word := range words {
		for _, char := range word {
			if byte(char) == x {
				ans = append(ans, i)
				break
			}
		}
	}
	return ans
}
