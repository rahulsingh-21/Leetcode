package main

func minSwaps(nums []int) int {
	to, te := 0, 0
	var odd, even []int

	for _, num := range nums {
		if num%2 == 0 {
			te++
			even = append(even, num)
		} else {
			to++
			odd = append(odd, num)
		}
	}
	//n := len(nums)
	if absII(to, te) > 1 {
		return -1
	}
	if to > te {
		return swap(nums, odd, even, "odd")
	} else if to < te {
		return swap(nums, odd, even, "even")
	} else {
		return min(swap(nums, odd, even, "odd"), swap(nums, odd, even, "even"))
	}
	return -1

}

func swap(nums, odd, even []int, seq string) int {
	swaps := 0
	p := 0
	switch seq {
	case "odd":
		for i := 0; i < len(nums); i++ {
			if nums[i]%2 != 0 {
				swaps += absII(i, p)
				p += 2
			}

		}
	case "even":
		for i := 0; i < len(nums); i++ {
			if nums[i]%2 == 0 {
				swaps += absII(i, p)
				p += 2
			}
		}
	}
	return swaps
}

func absII(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 2 4 6 8 10 1 3 5 7
// 2 1 6

// 2 4 6 5 7
// 2 4 5 6 7
// 2 5 4 6 7
// 2 5 4 7 6
