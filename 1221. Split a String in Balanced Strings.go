package main

func balancedStringSplit(s string) int {
	ans := 0
	count := 0

	for _, char := range s {
		if char == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			ans++
		}
	}
	return ans
}
