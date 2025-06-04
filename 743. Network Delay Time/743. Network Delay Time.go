package networkdelay

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {

	adj := make(map[int][]Node)

	for _, time := range times {
		src := time[0]
		adj[src] = append(adj[src], Node{time[1], time[2]})
	}

	distance := make([]int, n+1)

	for i := 0; i <= n; i++ {
		distance[i] = math.MaxInt32
	}

	distance[k] = 0

	minheap := &MinHeap{}
	heap.Push(minheap, Node{k, 0})

	for minheap.Len() > 0 {
		node := heap.Pop(minheap).(Node)

		for _, next := range adj[node.Index] {
			next.Cost += node.Cost
			if next.Cost < distance[next.Index] {
				heap.Push(minheap, next)
				distance[next.Index] = next.Cost
			}
		}
	}

	var ans int

	for i := 1; i <= n; i++ {
		if distance[i] == math.MaxInt32 {
			return -1
		}
		if distance[i] > ans {
			ans = distance[i]
		}
	}
	return ans
}

type Node struct {
	Index, Cost int
}

type MinHeap []Node

func (m MinHeap) Len() int            { return len(m) }
func (m MinHeap) Less(i, j int) bool  { return m[i].Cost < m[j].Cost }
func (m MinHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *MinHeap) Push(a interface{}) { *m = append(*m, a.(Node)) }
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	res := old[n-1]
	*m = old[:n-1]
	return res
}
