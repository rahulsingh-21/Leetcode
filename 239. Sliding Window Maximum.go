package main

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	var ans []int
	var deque []int

	for i := 0; i < len(nums); i++ {

		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}

	}

	return ans

}
