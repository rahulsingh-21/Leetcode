package main

import "container/heap"

func swimInWater(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])

	visited := make([][]bool, n)

	for i := range visited {
		visited[i] = make([]bool, m)
	}

	h := &minHeapII{}
	heap.Push(h, nodeI{0, 0, heights[0][0]})

	for h.Len() > 0 {
		height := heap.Pop(h).(nodeI)
		i, j, e := height.a, height.b, height.time

		if i == n-1 && j == m-1 {
			return e
		}

		if visited[i][j] {
			continue
		}

		visited[i][j] = true
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if x >= 0 && y >= 0 && x < n && y < m && !visited[x][y] {
				newTime := max(heights[x][y], e)
				heap.Push(h, nodeI{x, y, newTime})
			}
		}
	}
	return -1

}

type nodeI struct {
	a, b, time int
}

type minHeapII []nodeI

func (m minHeapII) Len() int           { return len(m) }
func (m minHeapII) Less(i, j int) bool { return m[i].time < m[j].time }
func (m minHeapII) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeapII) Push(x interface{}) {
	*m = append(*m, x.(nodeI))
}
func (m *minHeapII) Pop() interface{} {
	temp := *m
	last := temp[len(temp)-1]
	*m = temp[:len(temp)-1]
	return last
}
