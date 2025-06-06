package main

func asteroidCollision(asteroids []int) []int {

	var stack []int

	for i := 0; i < len(asteroids); i++ {

		if len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, asteroids[i])
		} else {

			top := stack[len(stack)-1]
			if asteroids[i] == -top {
				stack = stack[:len(stack)-1]
			} else if asteroids[i] < -top {
				stack = stack[:len(stack)-1]
				i--
			} else if asteroids[i] > 0 {
				stack = append(stack, asteroids[i])
			}
		}
	}
	return stack
}
