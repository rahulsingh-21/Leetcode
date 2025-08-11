package main

func maxPoints(points [][]int) int {
	ans := 0

	if len(points) == 1 {
		return 1
	}

	for i := 0; i < len(points); i++ {
		m := make(map[float64]int)
		for j := 0; j < len(points); j++ {
			if i != j {
				x1, x2, y1, y2 := points[i][0], points[j][0], points[i][1], points[j][1]
				slope := float64(y2-y1) / float64(x2-x1)

				m[slope]++
			}
		}
		for _, val := range m {
			ans = max(ans, val+1)
		}
	}

	return ans

}
