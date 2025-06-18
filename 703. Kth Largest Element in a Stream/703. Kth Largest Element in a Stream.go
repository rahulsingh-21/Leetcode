package main

import "container/heap"

type KthLargest struct {
	h *MinHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	return (*this.h)[0]
}

type MinHeap []int

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i] < m[j] }
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
