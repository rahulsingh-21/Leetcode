package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	temp := 1.0

	for n != 0 {
		if n&1 == 1 {
			temp *= x
		}
		x *= x
		n = n >> 1
	}
	return temp
}
