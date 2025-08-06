package main

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeapII{}

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

type MinHeapII []*ListNode

func (h MinHeapII) Len() int            { return len(h) }
func (h MinHeapII) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeapII) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeapII) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeapII) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
