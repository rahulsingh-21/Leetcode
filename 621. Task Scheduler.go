package main

import "container/heap"

func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)

	for _, task := range tasks {
		freq[task]++
	}

	h := &maxHeap{}

	for _, count := range freq {
		heap.Push(h, count)
	}

	count := 0
	for h.Len() > 0 {
		var temp []int

		for i := 0; i <= n; i++ {
			if h.Len() > 0 {
				val := heap.Pop(h).(int)
				val--
				if val > 0 {
					temp = append(temp, val)
				}
			}
			count++
			if h.Len() == 0 && len(temp) == 0 {
				break
			}
		}

		for _, val := range temp {
			heap.Push(h, val)
		}
	}
	return count
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
