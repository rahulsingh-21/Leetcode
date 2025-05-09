package main

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	m := map[[26]int][]string{}
	for _, val := range strs {
		key := [26]int{}
		for _, k := range val {
			key[k-'a']++
		}
		m[key] = append(m[key], val)
	}

	for _, keys := range m {
		ans = append(ans, keys)
	}
	return ans
}
