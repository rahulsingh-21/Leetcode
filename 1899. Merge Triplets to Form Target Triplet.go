package main

func mergeTriplets(triplets [][]int, target []int) bool {
	x, y, z := target[0], target[1], target[2]
	fx, fy, fz := false, false, false
	for _, t := range triplets {
		x1, y1, z1 := t[0], t[1], t[2]
		if x1 > x || y1 > y || z1 > z {
			continue
		}
		if x1 == x {
			fx = true
		}
		if y1 == y {
			fy = true
		}
		if z1 == z {
			fz = true
		}
		if fx && fy && fz {
			return true
		}
	}
	return false
}
