package main

func subsets(nums []int) [][]int {
	output := [][]int{{}}
	for _, num := range nums {
		var set [][]int
		for _, curr := range output {
			newSet := append([]int{}, curr...)
			newSet = append(newSet, num)
			set = append(set, newSet)
		}
		output = append(output, set...)
	}
	return output
}
