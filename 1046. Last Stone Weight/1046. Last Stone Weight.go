package main

import "container/heap"

func lastStoneWeight(stones []int) int {
	h := MinHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if x != y {
			heap.Push(&h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return (*&h)[0]
}

type MinHeap []int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *MinHeap) Pop() interface{} {
	temp := *m
	n := len(temp)
	last := temp[n-1]
	*m = temp[:n-1]
	return last
}
