package main

import "strconv"

func compress(chars []byte) int {

	idx := 0
	count := 1
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			count++
		} else {
			chars[idx] = chars[i]
			idx++
			if count > 1 {
				strArr := strconv.Itoa(count)
				for i := 0; i < len(strArr); i++ {
					chars[idx] = strArr[i]
					idx++
				}
				count = 1
			}
		}
	}

	chars[idx] = chars[len(chars)-1]
	idx++
	if count > 1 {
		strArr := strconv.Itoa(count)
		for i := 0; i < len(strArr); i++ {
			chars[idx] = strArr[i]
			idx++
		}
		count = 1
	}

	return idx
}
