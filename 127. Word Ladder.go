package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := map[string]struct{}{}

	q := []string{}
	q = append(q, beginWord)

	for _, word := range wordList {
		set[word] = struct{}{}
	}

	if _, ok := set[endWord]; !ok {
		return 0
	}

	count := 1
	for len(q) > 0 {
		count++
		for _, word := range q {

			for i, _ := range word {
				temp := []rune(word)
				for j := 'a'; j <= 'z'; j++ {
					temp[i] = j

					str := string(temp)

					if str == endWord {
						return count
					}

					if _, ok := set[str]; ok {
						q = append(q, str)
						delete(set, str)
					}
				}
			}
			q = q[1:]

		}

	}
	return 0
}
