package main

import "math"

func minWindow(s string, t string) string {

	target := make(map[byte]int)
	window := make(map[byte]int)

	for i := range t {
		target[t[i]-'a']++
	}

	have, need := 0, len(target)
	l := 0
	res := []int{-1, -1}
	reslen := math.MaxInt32
	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++

		if target[s[i]-'a'] > 0 && target[s[i]-'a'] == window[s[i]-'a'] {
			have++
		}

		for have == need {
			if i-l+1 < reslen {
				reslen = i - l + 1
				res = []int{l, i}
			}

			window[s[l]-'a']--
			if target[s[l]-'a'] > 0 && window[s[l]-'a'] < target[s[l]-'a'] {
				have--
			}
			l++
		}
	}

	if reslen == math.MaxInt32 {
		return ""
	}

	return s[res[0] : res[1]+1]

}
