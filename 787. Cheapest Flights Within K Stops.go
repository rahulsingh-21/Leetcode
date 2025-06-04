package main

import (
	"container/heap"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := make([][]Node, n)

	for _, flight := range flights {
		src := flight[0]
		adj[src] = append(adj[src], Node{flight[1], flight[2], 0})
	}

	minheap := &MinHeap{}
	heap.Push(minheap, Node{src, 0, 0})
	res := math.MaxInt64

	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt64
	}

	for minheap.Len() > 0 {
		node := heap.Pop(minheap).(Node)

		if node.Index == dst {
			res = min(res, node.Cost)
		}

		if node.Stops <= k {
			for _, next := range adj[node.Index] {
				next.Stops = node.Stops + 1
				next.Cost += node.Cost
				if next.Cost < distance[next.Index] {
					heap.Push(minheap, next)
					distance[next.Index] = next.Cost
				}
			}
		}
	}

	if res == math.MaxInt64 {
		return -1
	}

	return res

}

type Node struct {
	Index int
	Cost  int
	Stops int
}

type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i int, j int) bool {
	if h[i].Stops == h[j].Stops {
		return h[i].Cost < h[j].Cost
	}
	return h[i].Stops < h[j].Stops
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(a interface{}) {
	*h = append(*h, a.(Node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
