package main

func twoSum(numbers []int, target int) []int {
	ans := []int{}

	i := 0
	j := len(numbers) - 1

	for i < j {
		if numbers[i]+numbers[j] == target {
			ans = append(ans, i+1, j+1)
			return ans
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return ans
}
