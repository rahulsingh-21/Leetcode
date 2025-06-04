package main

func moveZeroes(nums []int) {
	i := 0
	for _, n := range nums {
		if n != 0 {
			nums[i] = n
			i++
		}
	}

	for j := i; j < len(nums); j++ {
		nums[j] = 0
	}
}
