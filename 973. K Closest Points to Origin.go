package main

import (
	"container/heap"
	"math"
)

func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)

	for _, point := range points {
		dis := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		heap.Push(h, Points{dis, point[0], point[1]})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res [][]int
	for i := 0; i < k; i++ {
		temp := heap.Pop(h).(Points)
		res = append(res, []int{temp.x, temp.y})
	}
	return res
}

type Points struct {
	distance float64
	x, y     int
}

type PointHeap []Points

func (p PointHeap) Len() int           { return len(p) }
func (p PointHeap) Less(i, j int) bool { return p[i].distance > p[j].distance }
func (p PointHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PointHeap) Push(v interface{}) {
	*p = append(*p, v.(Points))
}
func (p *PointHeap) Pop() interface{} {
	temp := *p
	n := len(temp)
	last := temp[n-1]
	*p = temp[:n-1]
	return last
}
