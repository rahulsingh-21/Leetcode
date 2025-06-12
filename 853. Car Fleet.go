package main

import "sort"

func carFleet(target int, position []int, speed []int) int {

	var cars []Cars

	for i := 0; i < len(speed); i++ {
		cars = append(cars, Cars{position[i], speed[i]})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	time := 0.0
	count := 0

	for _, car := range cars {
		distanceToReach := float64(target - car.position)
		timeToReach := distanceToReach / float64(car.speed)
		if time < timeToReach {
			time = timeToReach
			count++
		}
	}
	return count
}

type Cars struct {
	position, speed int
}
