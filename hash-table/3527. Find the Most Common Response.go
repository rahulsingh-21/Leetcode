package hashtable

import "sort"

func findCommonResponse(responses [][]string) string {
	m := map[string]int{}
	for i := 0; i < len(responses); i++ {
		set := map[string]struct{}{}
		for j := 0; j < len(responses[i]); j++ {
			set[responses[i][j]] = struct{}{}
		}
		for key, _ := range set {
			m[key]++
		}
	}
	max := 0
	ans := ""
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//fmt.Println(keys, m)
	for _, key := range keys {
		if m[key] > max {
			max = m[key]
			ans = key
		}
	}
	return ans
}
