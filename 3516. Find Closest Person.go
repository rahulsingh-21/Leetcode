package math

import "fmt"

func findClosest(x int, y int, z int) int {
	var n1, n2 int
	if z > x {
		n1 = z - x
	} else {
		n1 = x - z
	}

	if z > y {
		n2 = z - y
	} else {
		n2 = y - z
	}
	fmt.Println(n1, n2)
	if n1 < n2 {
		return 1
	} else if n1 > n2 {
		return 2
	} else {
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
