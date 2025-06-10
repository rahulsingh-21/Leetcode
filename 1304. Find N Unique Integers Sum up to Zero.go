package main

func sumZero(n int) []int {

	x := n / 2
	m := 1

	var ans []int

	for i := 0; i < x; i++ {
		ans = append(ans, m, -m)
		m++
	}

	if n%2 == 0 {
		return ans
	}

	ans = append(ans, 0)
	return ans
}
