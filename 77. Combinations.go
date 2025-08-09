package main

func combine(n int, k int) [][]int {
	ans := [][]int{}

	var backtrack func(i int, curr []int)
	backtrack = func(i int, curr []int) {
		if len(curr) == k {
			temp := append([]int{}, curr...)
			ans = append(ans, temp)
			return
		}

		for j := i; j <= n; j++ {
			curr = append(curr, j)
			backtrack(j+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(1, []int{})
	return ans

}
