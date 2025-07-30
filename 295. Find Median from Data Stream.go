package main

import "container/heap"

type MedianFinder struct {
	left  *MaxHeap
	right *minHeap
}

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func ConstructorXV() MedianFinder {
	return MedianFinder{&MaxHeap{}, &minHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.left, num)

	heap.Push(this.right, heap.Pop(this.left))

	if this.left.Len() < this.right.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() > this.right.Len() {
		return float64((*this.left)[0])
	}

	return (float64((*this.left)[0]) + float64((*this.right)[0])) / 2.0
}
