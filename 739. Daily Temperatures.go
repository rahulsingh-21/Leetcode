package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)

	for i, _ := range temperatures {
		ans[i] = -1
	}

	st := make([]int, n)
	res := make([]int, n)
	for i := 0; i < n; i++ {

		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			res[idx] = i - idx
		}
		st = append(st, i)
	}
	return res

}
