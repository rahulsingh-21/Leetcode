package boyermoore

func majorityElement(nums []int) []int {
	n1, n2, c1, c2 := 0, 0, 0, 0

	for _, val := range nums {
		if n1 == val {
			c1++
		} else if n2 == val {
			c2++
		} else if c1 == 0 {
			n1 = val
			c1++
		} else if c2 == 0 {
			n2 = val
			c2++
		} else {
			c1--
			c2--
		}
	}

	var ans []int
	c1, c2 = len(nums)/3, len(nums)/3

	for _, val := range nums {
		if val == n1 {
			c1--
		} else if val == n2 {
			c2--
		}
	}

	if c1 < 0 {
		ans = append(ans, n1)
	}
	if c2 < 0 {
		ans = append(ans, n2)
	}
	return ans
}
