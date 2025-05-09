package main

func isHappy(n int) bool {
	m := map[int]int{}

	for {
		temp := getNo(n)
		m[temp]++
		if temp == 1 {
			return true
		}
		if m[temp] > 1 {
			return false
		}
		n = temp
	}
	return false
}

func getNo(n int) int {
	ans := 0
	for ; n > 0; n /= 10 {
		ans += (n % 10) * (n % 10)
	}
	return ans
}
