package main

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode

	secondHalf := slow.Next
	slow.Next = nil
	for secondHalf != nil {
		temp := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = temp
	}

	left := head

	for prev != nil {
		temp1 := left.Next
		temp2 := prev.Next
		left.Next = prev
		prev.Next = temp1
		left = temp1
		prev = temp2

	}

}
