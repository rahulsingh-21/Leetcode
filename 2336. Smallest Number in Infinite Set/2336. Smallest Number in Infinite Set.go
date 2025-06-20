package main

import "container/heap"

type SmallestInfiniteSet struct {
	h   *MinHeap
	set map[int]struct{}
	n   int
}

func Constructor() SmallestInfiniteSet {
	h := &MinHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{h, make(map[int]struct{}), 1}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.h.Len() > 0 {
		temp := heap.Pop(this.h).(int)
		delete(this.set, temp)
		return temp
	}
	smallest := this.n
	this.n++
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.set[num]; !ok && num < this.n {
		this.set[num] = struct{}{}
		heap.Push(this.h, num)
	}
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
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
