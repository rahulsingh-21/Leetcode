package main

func canCompleteCircuit(gas []int, cost []int) int {
	sum, maxi := 0, 0
	ans := -1
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		maxi += gas[i] - cost[i]
		if sum < 0 {

			sum = 0
			ans = i + 1
		}

	}
	if maxi >= 0 {
		return ans
	}
	return -1
}
