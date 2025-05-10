package main

func maxFreqSum(s string) int {
	freq := [26]int{}
	for _, val := range s {
		freq[val-'a']++
	}
	mv, mc := 0, 0
	for i, val := range freq {
		if i == 'a'-'a' || i == 'e'-'a' || i == 'i'-'a' || i == 'o'-'a' || i == 'u'-'a' {
			mv = max(mv, val)
		} else {
			mc = max(mc, val)
		}
	}
	//fmt.Println(mc,mv)
	return mc + mv
}
