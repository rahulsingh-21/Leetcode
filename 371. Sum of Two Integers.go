package main

func getSum(a int, b int) int {
	sum := a ^ b
	ca := a & b
	for ca != 0 {
		ca = ca << 1
		sum, ca = sum^ca, sum&ca
	}
	return sum
}
