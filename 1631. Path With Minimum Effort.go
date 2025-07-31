package main

import (
	"container/heap"
	"math"
)

func minimumEffortPath(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])

	distance := make([][]int, n)

	for i := range distance {
		distance[i] = make([]int, m)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}

	distance[0][0] = 0
	h := &minHeapI{}
	heap.Push(h, node{0, 0, 0})

	for h.Len() > 0 {
		height := heap.Pop(h).(node)
		i, j, e := height.a, height.b, height.effort

		if i == n-1 && j == m-1 {
			return e
		}

		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if x >= 0 && y >= 0 && x < n && y < m {
				newEffort := max(abs(heights[x][y]-heights[i][j]), e)
				if newEffort < distance[x][y] {
					distance[x][y] = newEffort
					heap.Push(h, node{x, y, newEffort})
				}
			}
		}
	}
	return 0

}

type node struct {
	a, b, effort int
}

type minHeapI []node

func (m minHeapI) Len() int           { return len(m) }
func (m minHeapI) Less(i, j int) bool { return m[i].effort < m[j].effort }
func (m minHeapI) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeapI) Push(x interface{}) {
	*m = append(*m, x.(node))
}
func (m *minHeapI) Pop() interface{} {
	temp := *m
	last := temp[len(temp)-1]
	*m = temp[:len(temp)-1]
	return last
}
