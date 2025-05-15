package main

func wordSubsets(words1 []string, words2 []string) []string {
	freq := make([]int, 26)

	for _, word := range words2 {
		temp := getFreq(word)
		for key, value := range freq {
			if temp[key] > value {
				freq[key] = temp[key]
			}
		}
	}

	var ans []string
	for _, word := range words1 {
		temp := getFreq(word)
		flag := true
		for key, value := range freq {
			if temp[key] < value {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, word)
		}

	}

	return ans

}

func getFreq(s string) []int {
	m := make([]int, 26)
	for _, str := range s {
		m[str-'a']++
	}
	return m
}
