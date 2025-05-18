package main

func smallestIndex(nums []int) int {
	for i, n := range nums {
		temp := sum(n)
		if i == temp {
			return i
		}
	}
	return -1
}

func sum(n int) int {
	ans := 0
	for i := n; i > 0; i /= 10 {
		ans += i % 10

	}
	return ans
}
