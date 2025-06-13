package main

type NodeII struct {
	Val    int
	Next   *NodeII
	Random *NodeII
}

func copyRandomList(head *NodeII) *NodeII {
	curr := head

	m := make(map[*NodeII]*NodeII)

	for curr != nil {
		m[curr] = &NodeII{Val: curr.Val}
		curr = curr.Next

	}

	curr = head
	for curr != nil {
		m[curr].Next = m[curr.Next]
		m[curr].Random = m[curr.Random]
		curr = curr.Next
	}
	return m[head]
}
